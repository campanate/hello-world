def label = "worker-${UUID.randomUUID().toString()}"

podTemplate(label: label, serviceAccount: "jenkins", containers: [
  containerTemplate(name: 'docker', image: 'docker:dind', command: 'cat', ttyEnabled: true),
],
resourceRequestCpu: '2048m',
resourceLimitCpu: '4096m',
resourceRequestMemory: '2048Mi',
resourceLimitMemory: '4096Mi',
volumes: [
  hostPathVolume(mountPath: '/var/run/docker.sock', hostPath: '/var/run/docker.sock'),
]) {
  node(label) {
    // Handle variables
    def accountNumber = ""
    def localImageName = "hello-world"
  

    // Git variables

    def myRepo = checkout scm
    def gitBranch = myRepo.GIT_BRANCH
    def gitCommit = myRepo.GIT_COMMIT
    def dockerUsername = "vcampanate"
   
    stage('Build and Test image') {
      container('docker') {
          withCredentials([[$class: 'UsernamePasswordMultiBinding',
          credentialsId: 'dockerhub',
          usernameVariable: 'DOCKER_USERNAME',
          passwordVariable: 'DOCKER_PASSWORD']]) {
            sh """
                docker login -u ${dockerUsername} -p '${DOCKER_PASSWORD}'
                docker build --network=host -t docker.io/${dockerUsername}/${localImageName}:${gitCommit} .
                docker run --name hello-world-test docker.io/${dockerUsername}/${localImageName}:${gitCommit} go test
                docker push docker.io/${dockerUsername}/${localImageName}:${gitCommit}
                docker rmi -f docker.io/${dockerUsername}/${localImageName}:${gitCommit}
              """
            dockerUsername = "${dockerUsername}"

          }
      }
    }

    stage("Put image on property files") {
      writeFile file: 'image.properties', text: "repository=docker.io/${dockerUsername}/${localImageName}\ntag=${gitCommit}"
      archiveArtifacts artifacts: 'image.properties', excludes: ''
      
    }
  
  }

}