package parser

import (
  "fmt"
  "strings"
)

//go:generate make generate

%%{
  machine phone;

  # sets the start point for the current value
  action mark{ m.pb = m.p }

  # writes international code.
  action set_intcode{
    phone.IntCode = m.string()
  }

  # writes area code
  action set_area{
    phone.AreaCode = m.string()
  }

  # return area-related error
  action err_area{
    m.err = fmt.Errorf("invalid area code, expected 200..999")
  }

  # writes number and replaces all hyphens and spaces.
  action set_number{
    n := strings.ReplaceAll(m.string(), "-","")
    phone.Number = strings.ReplaceAll(n, " ","")
  }

  # return an error when phone with incorrect format.
  action err_phone{
    m.err = fmt.Errorf("invalid phone format: %s", m.data)
  }

  # sets default international code.
  action set_defaults{
    if phone.IntCode == ""{
      phone.IntCode = "1"
    }
  }

  sp = ' ';
  hy = (sp* | '-');
  hbl = ( hy | '(' );
  hbr = ( hy | ')' );

  # defines an international code: +1,1
  int = '+'? '1' >mark %set_intcode;

  # defines a range 200..999
  nxx = ('2'..'9'  . digit{2});

  # defines an area code with separators: (NXX), -NXX-
  area = hbl? nxx >mark %set_area @err(err_area) hbr?;

  # defines phone number in range 2000000..9999999, with an optional separators.
  number =  (nxx hy? digit{2} hy? digit{2}) >mark %set_number @err(err_phone);

  # defines main machine for a phone.
  main := int? sp* area sp* number @set_defaults;

  write data;

	access m.;

  variable p m.p;
  variable pe m.pe;
  variable eof m.eof;
  variable data m.data;

}%%

// Machine contains variables used by Ragel auto generated code.
// See Ragel docs for details.
type Machine struct {
  // Data to process.
  data                []byte
  // Data pointer.
  p int
  // Data end pointer.
  pe int
  // Curent state.
  cs  int
  // End of file pointer.
  eof int
  // Start of current date block.
  pb  int
  // Current err
  err error
}

// Phone represents parser result and will be filled by Raagel actions.
type Phone struct {
  IntCode   string
  AreaCode  string
  Number    string
}

// New initialized new Machine structure.
func New() *Machine {
  return &Machine{}
}

// string returns current parsed variable. Variable m.pb should be updated by
// "mark" action, while m.p is a current parser position inside m.data.
func (m *Machine) string() string {
  return string(m.data[m.pb:m.p])
}

// Parse takes a slice of bytes as an input an fills Phone structure in case
// input is a phone number in one of valid formats.
func (m *Machine) Parse(input []byte) (*Phone, error) {
  // Initialize variables required by Ragel.
  m.data = input
  m.pe = len(input)
  m.p = 0
  m.pb = 0
  m.err = nil
	m.eof = len(input)

  phone := &Phone{}

  %% write init;
  %% write exec;

  if m.err != nil {
      return nil, m.err
  }

  return phone, m.err
}

