pipeline {
    agent any
    stages {
        stage('Pre test') {
            steps {
                echo 'Pre test'
                echo ${BUILD_ID}
            }
        }

        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'make build BUILD_TAG=${BUILD_ID}'
            }
        }
    }
}