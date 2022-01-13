pipeline {
    agent any
    environment {
        BUILD_TAG='${BUILD_ID}'
    }
    options {
        retry(3)
    }
    stages {
        stage('INFO') {
            steps {
                sh 'echo "Job:      ${JOB_NAME}"'
                sh 'echo "Build id: ${BUILD_ID}"'
            }
        }

        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'make build'
            }
        }

        stage('Test') {
            steps {
                echo 'Testing...'
                sh 'make test'
            }
        }

        stage('Build docker image') {
            steps {
                echo 'Building docker image...'
                sh 'make build-docker-image'
            }
        }

        stage('Push docker image') {
            steps {
                echo 'Pushing to docker images repo...'
                sh 'make push-docker-image'
            }
        }
    }
}