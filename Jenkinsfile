pipeline {
    agent any
    tools {
        go 'Go 1.17'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Build Hello') {
            steps {
                echo 'Building...'
                sh 'go build .'
                archiveArtifacts artifacts: 'go_jenkins', fingerprint: true
            }
        }

        stage('Testing') {
            steps {
                echo 'Testing...'
                sh 'go test'
            }
        }

        stage('Benchmark') {
            steps {
                echo 'Benchmarking...'
                sh 'go test -bench .'
            }
        }

        stage('Deploy bin') {
            steps {
                echo 'Deploying'
                sh 'ls -l'
            }
        }
    }
}