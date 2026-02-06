package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Student struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

var fileName = "students.json"

func readStudents() ([]Student, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		// Return empty slice if file doesn't exist
		if os.IsNotExist(err) {
			return []Student{}, nil
		}
		return nil, err
	}

	var students []Student
	json.Unmarshal(data, &students)
	return students, nil
}

func writeStudents(students []Student) error {
	data, err := json.MarshalIndent(students, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

// ---------- HANDLERS ----------

// GET /students
func getStudents(w http.ResponseWriter, _ *http.Request) {
	students, _ := readStudents()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

// POST /students
func createStudent(w http.ResponseWriter, r *http.Request) {
	var newStudent Student
	json.NewDecoder(r.Body).Decode(&newStudent)

	students, _ := readStudents()

	newStudent.ID = int64(len(students) + 1)
	students = append(students, newStudent)

	writeStudents(students)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newStudent)
}

// PUT /students?id=1
func updateStudent(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	var updated Student
	json.NewDecoder(r.Body).Decode(&updated)

	students, _ := readStudents()

	for i := range students {
		if students[i].ID == id {
			students[i] = updated
			students[i].ID = id
		}
	}

	writeStudents(students)
	json.NewEncoder(w).Encode(updated)
}

// PATCH /students?id=1
func patchStudent(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	body, _ := io.ReadAll(r.Body)

	var partial map[string]interface{}
	json.Unmarshal(body, &partial)

	students, _ := readStudents()

	for i := range students {
		if students[i].ID == id {
			if name, ok := partial["name"].(string); ok {
				students[i].Name = name
			}
			if age, ok := partial["age"].(float64); ok {
				students[i].Age = int64(age)
			}
		}
	}

	writeStudents(students)
	json.NewEncoder(w).Encode(students)
}

// DELETE /students?id=1
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	students, _ := readStudents()

	var updated []Student
	for _, s := range students {
		if s.ID != id {
			updated = append(updated, s)
		}
	}

	writeStudents(updated)
	w.Write([]byte("Deleted"))
}

// ---------- ROUTER ----------

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		getStudents(w, r)

	case http.MethodPost:
		createStudent(w, r)

	case http.MethodPut:
		updateStudent(w, r)

	case http.MethodPatch:
		patchStudent(w, r)

	case http.MethodDelete:
		deleteStudent(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/students", studentsHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
