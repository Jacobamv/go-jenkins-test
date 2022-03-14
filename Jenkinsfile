pipeline {
    agent any
    tools {
        go 'go-1.17.8'
    }
    environment {
        GO1178MODULE = 'on'
    }
    stages {
        stage('Compile') {
            steps {
                sh 'cd cmd/ && go build main.go'
            }
        }
    }
}