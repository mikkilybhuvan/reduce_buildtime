pipeline {
    agent any

    environment {
        GOROOT = "C:\\Program Files\\Go"
        PATH = "${GOROOT}\\bin;${env.PATH}"
    }

    stages {
        stage('Verify Go') {
            steps {
                bat 'where go'
                bat 'go version'
            }
        }

        stage('Build App') {
            steps {
                echo "Building main app..."
                bat 'go build -o bin\\myapp.exe ./cmd/myapp'
    }
}


        stage('Build Test Client') {
            when {
                expression { fileExists('go_testclient.go') }
            }
            steps {
                echo "Building test client..."
                bat 'go build -o bin\\testclient.exe go_testclient.go'
            }
        }

        stage('Run Tests') {
            steps {
                echo "Running Go unit tests..."
                bat 'go test ./...'
            }
        }
    }

    post {
        success {
            echo "Build completed successfully!"
        }
        failure {
            echo "Build failed. Check logs above."
        }
    }
}