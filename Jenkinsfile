pipeline {
    agent any
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