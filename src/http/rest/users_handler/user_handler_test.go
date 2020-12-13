package users_handler

// func TestCreateUser(t *testing.T) {
// 	client := &http.Client{}
// 	payLoad := []byte(`{"user_name": "tim", "email": "timkotowski@gmail.com", "password": "msdasdjkj3213"}`)
// 	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/create", bytes.NewBuffer(payLoad))
// 	if err != nil {
// 		t.Errorf("Expect nil, recived %s", err.Error())
// 	}
// 	res, err := client.Do(req)
// 	if err != nil {
// 		t.Errorf("Expect nil, recived %s", err.Error())
// 	}
// 	defer res.Body.Close()

// 	var user users.User
// 	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
// 		 t.Errorf("Expect nil, recived %s", err.Error())
// 	}
// 	if user.ID == 0 {
// 		t.Errorf("Expected id to not be nil, got '%v'", user.ID)
// 	}
// 	// Check if username is correct.
// 	if user.Username != "tim" {
// 		t.Errorf("Expected username to be 'tim', got '%s'", user.Username)
// 	}
// 	if user.Email != "timkotowski@gmail.com" {
// 		t.Errorf("Expected username to be 'timkotowski@gmail.com', got '%s'", user.Username)

// 	}
// 	fmt.Println(user)
// }

// func TestGetUser(t *testing.T) {
// 	client := &http.Client{}
// 	payLoad := []byte(`{"user_name": "johnnybravo", "email": "johnnybravo@gmail.com", "password": "johnnyBravo12345!"}`)
// 	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/user", bytes.NewBuffer(payLoad))
// 	if err != nil {
// 		t.Errorf("Expect nil, recived %s", err.Error())
// 	}
// 	res, err := client.Do(req)
// 	if err != nil {
// 		t.Errorf("Expect nil, recived %s", err.Error())
// 	}
// 	defer res.Body.Close()

// 	var user users.User
// 	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
// 		 t.Errorf("Expect nil, recived %s", err.Error())
// 	}
// 	if user.ID == 0 {
// 		t.Errorf("Expected id to not be nil, got '%v'", user.ID)
// 	}
// 	// Check if username is correct.
// 	if user.Username != "johnnybravo" {
// 		t.Errorf("Expected username to be 'johnnybravo', got '%s'", user.Username)
// 	}
// 	if user.Email != "johnnybravo@gmail.com" {
// 		t.Errorf("Expected username to be 'johnnybravo@gmail.com', got '%s'", user.Username)
// 	}
// 	fmt.Println(user)
// }
