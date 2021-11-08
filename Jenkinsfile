pipeline {
    agent any
    tools {
        go 'go-1.17'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Build Hello') {
            steps {
                echo 'Building...'
                sh 'go build .'
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