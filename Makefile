VERSION=0.1.0

dist: all
	mkdir -p dist/todos
	cp ./todos-backend-server/build/linux/todos_backend_server ./dist/todos/
	cp ./todos-backend-server/build/win/todos_backend_server.exe ./dist/todos/
	cp -R ./todos-frontend/build ./dist/todos/www
	cp README.md ./dist/todos/
#  cp CHANGELOG.md ./dist/todos/
	tar czf ./dist/todos-v${VERSION}.tar.gz -C ./dist ./todos

all:
	${MAKE} -C todos-backend-server all
	cd todos-frontend && CI=true npm run all

run: build
	${MAKE} -C todos-backend-server run

build:
	${MAKE} -C todos-backend-server build
	cd todos-frontend && CI=true npm run build

clean:
	${MAKE} -C todos-backend-server clean
	cd todos-frontend && npm run clean
