package main

type Material struct {
	ID          int
	Name        string
	Quantity    int
	Description string
}

// GetAllMaterials listará los materiales en inventario
func GetAllMaterials() ([]Material, error) {
	rows, err := DB.Query("SELECT id, name, quantity, description FROM materials")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var materials []Material
	for rows.Next() {
		var m Material
		rows.Scan(&m.ID, &m.Name, &m.Quantity, &m.Description)
		materials = append(materials, m)
	}
	return materials, nil
}

// GetMaterialByID buscará materiales por ID
func GetMaterialByID(id int) (Material, error) {
	var m Material
	err := DB.QueryRow("SELECT id, name, quantity, description FROM materials WHERE id=$1", id).Scan(
		&m.ID, &m.Name, &m.Quantity, &m.Description)
	return m, err
}

// CreateMaterial agregará un nuevo material al inventario
func CreateMaterial(m Material) error {
	_, err := DB.Exec("INSERT INTO materials (name, quantity, description) VALUES ($1, $2, $3)",
		m.Name, m.Quantity, m.Description)
	return err
}

// UpdateMaterial actualiza el inventario según cambios realizados
func UpdateMaterial(m Material) error {
	_, err := DB.Exec("UPDATE materials SET name=$1, quantity=$2, description=$3 WHERE id=$4",
		m.Name, m.Quantity, m.Description, m.ID)
	return err
}

// DeleteMaterial elimina algún material del inventario
func DeleteMaterial(id int) error {
	_, err := DB.Exec("DELETE FROM materials WHERE id=$1", id)
	return err
}
