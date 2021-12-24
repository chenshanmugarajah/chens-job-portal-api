package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/chenshanmugarajah/chens-job-portal-api/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://chen:Danish24@cluster0.jptci.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "uk-job-market"
const colName = "jobs"

//MOST IMPORTANT
var collection *mongo.Collection

// connect with monogoDB

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
}

// MONGODB helpers - file

// insert 1 record
func insertOneJob(job models.Job) {
	inserted, err := collection.InsertOne(context.Background(), job)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

// update 1 record
func updateOneJob(jobId string, job models.Job) {
	id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"title": job.Title, "company": job.Company, "link": job.Link, "salary": job.Salary, "experience": job.Experience}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// get 1 job
func getOneJob(jobId string) primitive.M {
	id, _ := primitive.ObjectIDFromHex(jobId)
	var job bson.M
	err := collection.FindOne(
		context.Background(),
		bson.D{{"_id", id}},
	).Decode(&job)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		log.Fatal(err)
	}
	return job
}

// delete 1 record
func deleteOneJob(jobId string) {
	id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Job got delete with delete count: ", deleteCount)
}

// delete all records from mongodb
func deleteAllJobs() int64 {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("NUmber of movies delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all movies from database

func getAllJobs() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var jobs []primitive.M

	for cur.Next(context.Background()) {
		var job bson.M
		err := cur.Decode(&job)
		if err != nil {
			log.Fatal(err)
		}
		jobs = append(jobs, job)
	}

	defer cur.Close(context.Background())
	return jobs
}

// Actual controller - file

func GetMyAllJobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allJobs := getAllJobs()
	json.NewEncoder(w).Encode(allJobs)
}

func GetOneJobOnly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	job := getOneJob(params["id"])
	json.NewEncoder(w).Encode(job)
}

func CreateJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var job models.Job
	_ = json.NewDecoder(r.Body).Decode(&job)
	insertOneJob(job)
	json.NewEncoder(w).Encode(job)

}

func UpdateOneJobOnly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	var job models.Job
	_ = json.NewDecoder(r.Body).Decode(&job)
	params := mux.Vars(r)
	updateOneJob(params["id"], job)
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneJob(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllJobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllJobs()
	json.NewEncoder(w).Encode(count)
}
