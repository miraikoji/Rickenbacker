deploy:
	aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 812432336965.dkr.ecr.ap-northeast-1.amazonaws.com
	docker build -t rickenbacker -f Dockerfile.prod .
	docker tag rickenbacker:latest 812432336965.dkr.ecr.ap-northeast-1.amazonaws.com/rickenbacker:latest
	docker push 812432336965.dkr.ecr.ap-northeast-1.amazonaws.com/rickenbacker:latest
	aws ecs update-service --cluster rickenbacker-cluster --service rickenbacker-service --force-new-deployment
