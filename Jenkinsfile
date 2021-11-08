pipeline {
    agent any
    stages {
        stage('Build Hello') {
            steps {
                echo 'Building...'
                bat 'go build .'
                archiveArtifacts artifacts: '**/*.exe', fingerprint: true
            }
        }

        stage('Testing') {
            steps {
                echo 'Testing...'
            }
        }
    }
}