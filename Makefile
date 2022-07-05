VERSION=0.1.0

all: dist

build:
	${MAKE} -C todos-backend-server
	cd todos-frontend && CI=true npm run all

clean:
	${MAKE} -C todos-backend-server clean
	cd todos-frontend && npm run clean

dist: build
	mkdir -p dist/todos
	cp ./todos-backend-server/build/linux/todos_backend_server ./dist/todos/
	cp ./todos-backend-server/build/win/todos_backend_server.exe ./dist/todos/
	cp -R ./todos-frontend/build ./dist/todos/www
	cp README.md ./dist/todos/
#  cp CHANGELOG.md ./dist/todos/
	tar czf ./dist/todos-v${VERSION}.tar.gz -C ./dist ./todos
