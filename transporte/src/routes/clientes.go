
func Clientes_GetSchema() ([]models.Base_DB, string) {
	var Clientes []models.Base_DB
	tableName := "Fina_" + "Clientes"
	Clientes = append(Clientes, models.Base_DB{
		Name:        "c_docu",
		Description: "c_docu",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Min:  2,
			Max:  2,
		},
	})
	Clientes = append(Clientes, models.Base_DB{
		Name:        "n_docu",
		Description: "n_docu",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       1.100000,
			Max:       11,
			UpperCase: true,
		},
	})
	Clientes = append(Clientes, models.Base_DB{
		Name:        "l_clie",
		Description: "l_clie",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       10.000000,
			Max:       100,
			UpperCase: true,
		},
	})
	Clientes = append(Clientes, models.Base_DB{
		Name:        "k_gene",
		Description: "k_gene",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 10,
		},
	})
	Clientes = append(Clientes, models.Base_DB{
		Name:        "f_naci",
		Description: "f_naci",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Date: true,
		},
	})
	Clientes = append(Clientes, models.Base_DB{
		Name:        "l_dire",
		Description: "l_dire",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       40.000000,
			Max:       400,
			UpperCase: true,
		},
	})
	Clientes = append(Clientes, models.Base_DB{
		Name:        "l_refe",
		Description: "l_refe",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       40.000000,
			Max:       400,
			UpperCase: true,
		},
	})
	Clientes = append(Clientes, models.Base_DB{
		Name:        "c_ubig",
		Description: "c_ubig",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Min:  6,
			Max:  6,
		},
	})
	Clientes = append(Clientes, models.Base_DB{
		Name:        "n_tele",
		Description: "n_tele",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       3.000000,
			Max:       30,
			UpperCase: true,
		},
	})
	Clientes = append(Clientes, models.Base_DB{
		Name:        "n_celu",
		Description: "n_celu",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       3.000000,
			Max:       30,
			UpperCase: true,
		},
	})
	Clientes = append(Clientes, models.Base_DB{
		Name:        "l_obse",
		Description: "l_obse",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       10.000000,
			Max:       100,
			UpperCase: true,
		},
	})
	return Clientes, tableName
}
