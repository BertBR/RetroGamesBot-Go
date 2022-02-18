local:
	npx localtunnel --port 3000 --subdomain retrogames&
	
dev: local
	cd pkg && air