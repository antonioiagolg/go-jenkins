pipeline {
    agent any
    stages {
        stage('Build Hello') {
            steps {
                echo 'Building...'
                go build .
            }
        }

        stage('Testing') {
            steps {
                echo 'Testing...'
            }
        }
    }
}