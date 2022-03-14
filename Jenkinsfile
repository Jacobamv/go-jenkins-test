pipeline {
    agent any
    tools {
        go 'go 1.17.8'
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
        stage("Set a service"){
            steps{
                sh 'sudo cp server/go_jenkins_test.service /lib/systemd/system/go_jenkins_test.service'
                sh 'sudo service go_jenkins_test start'
            }
        }
        stage("Set a nginx"){
            steps{
                sh 'sudo cp server/go_jenkins_test /etc/nginx/sites-available/'
                sh 'sudo ln -sf /etc/nginx/sites-available/go_jenkins_test /etc/nginx/sites-enabled/'
                sh 'sudo service nginx restart'
           }
        }
    }
}
