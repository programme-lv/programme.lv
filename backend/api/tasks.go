package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/KrisjanisP/deikstra/service/models"
)

func (c *Controller) listTasks(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	var tasks []models.Task
	result := c.database.Find(&tasks)
	log.Println(result)

	// echo back the task
	resp, err := json.Marshal(tasks)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// send the response
	_, err = w.Write(resp)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func getTask(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (c *Controller) createTask(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	err := r.ParseMultipartForm(50 * (1 << 20)) // ~ 50 MB
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	mForm := r.MultipartForm
	for k := range mForm.File {
		file, fileHeader, err := r.FormFile(k)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}
		fmt.Printf("the uploaded file: name[%s], size[%d], header[%#v]\n",
			fileHeader.Filename, fileHeader.Size, fileHeader.Header)

		localFileName := "/srv/deikstra/tasks/" + fileHeader.Filename
		out, err := os.Create(localFileName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		_, err = io.Copy(out, file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		fmt.Printf("file %s uploaded ok\n", fileHeader.Filename)

		err = file.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		err = out.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}

	w.WriteHeader(200)
}

func (c *Controller) deleteTask(w http.ResponseWriter, r *http.Request) {
	// CORS
	log.Println(r.Body)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	// decode the request
	var task models.Task
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&task)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.database.Delete(&task).Error
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("deleted code: ", task.Code)
	log.Println("deleted name: ", task.Name)

	// echo back the task
	resp, err := json.Marshal(task)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// send the response
	_, err = w.Write(resp)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
