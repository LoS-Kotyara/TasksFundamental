pipeline {
    agent any
    environment {
        BUILD_TAG='${BUILD_ID}'
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
                retry(3) {
                    sh 'make build-docker-image'
                }
            }
        }

        stage('Push docker image') {
            steps {
                echo 'Pushing to docker images repo...'
                retry(2) {
                    sh 'make push-docker-image'
                }
            }
        }
    }
}