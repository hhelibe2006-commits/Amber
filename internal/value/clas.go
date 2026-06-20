package value

import "github.com/spf13/cobra"

type Clas struct {
	Typ    string
	Output string
	Input  PathList
	Cdc    string
}

func NewClas() *Clas {
	clas := new(Clas)
	return clas
}

func (clas *Clas) Set(cmd *cobra.Command) error {
	if typ, err := cmd.Flags().GetString("type"); err != nil {
		return err
	} else {
		clas.Typ = typ
	}
	return nil
}
