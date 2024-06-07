# Переменные
APP_NAME := MyProjectsWebSite
BUILD_DIR := ./build
SRC := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Компиляция приложения
build:
	@echo "Сборка приложения..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC)

# Запуск приложения
run: build
	@echo "Запуск приложения..."
	@$(BUILD_DIR)/$(APP_NAME)

# Очистка сборочных файлов
clean:
	@echo "Очистка..."
	@rm -rf $(BUILD_DIR)

# Перезапуск приложения после изменений
restart: clean run

.PHONY: build run clean restart
