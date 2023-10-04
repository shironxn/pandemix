# Pandemix

<img src="https://i.imgur.com/LKoI8s0.jpg" width="100%"></img>

The project, named "pandemix," is an exercise project that implements a REST API to manage patient data during a pandemic. Although this is just an exercise, the project provides an understanding of RESTful API concepts and their usage.

## Note

This project does not utilize a database, patient data is stored and managed within the application's memory.

## Usage

1. Clone this repository: `git clone <repository URL>`
2. Navigate to the project directory: `cd pandemix`
3. Install dependencies: `go mod tidy`
4. Run the built code: `go run main.go`

## API Reference

### Covid

| Method | Url          | Description       |
| :----- | :----------- | :---------------- |
| `GET`  | `/api/covid` | Get COVID-19 data |

### Patient

| Method | Url                  | Description                |
| :----- | :------------------- | :------------------------- |
| `GET`  | `/api/patient`       | Get all patient data.      |
| `GET`  | `/api/patient/${id}` | Get patient data by ID.    |
| `POST` | `/api/patient`       | Create new patient data.   |
| `PUT`  | `/api/patient/${id}` | Update patient data by ID. |
