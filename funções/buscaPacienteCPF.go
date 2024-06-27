func buscaPacienteCPF(cpf string) (Pacientes, error) {
    // Validar o CPF(mudar nome da função conforme o nome q a pessoa q fez colocou)
    if !validarCPF(cpf) {
        return nil, fmt.Errorf("CPF inválido: %s", cpf)
    }

    busca, err := db.Query(`SELECT * FROM paciente WHERE cpf LIKE concat('%', text($1), '%')`, cpf)
    if err != nil {
        return nil, err
    }
    defer busca.Close() 

    // Processar os resultados
    pacientes := make(Pacientes, 0)
    for busca.Next() {
        var paciente Paciente
        
        err := busca.Scan(&paciente.ID, &paciente.Nome, &paciente.CPF)
        
        if err != nil {
            return nil, err
        }

        pacientes = append(pacientes, paciente)
    }

    // Retornar os resultados
    return pacientes, nil
}
