#!/bin/bash

# RabbitMQ Docker management script

case "$1" in
    start)
        echo "Starting RabbitMQ..."
        docker-compose up -d
        echo "RabbitMQ is starting up..."
        echo "Management UI will be available at: http://localhost:15672"
        echo "Username: admin, Password: admin123"
        ;;
    stop)
        echo "Stopping RabbitMQ..."
        docker-compose down
        ;;
    restart)
        echo "Restarting RabbitMQ..."
        docker-compose restart
        ;;
    logs)
        docker-compose logs -f rabbitmq
        ;;
    status)
        docker-compose ps
        ;;
    *)
        echo "Usage: $0 {start|stop|restart|logs|status}"
        echo ""
        echo "Commands:"
        echo "  start   - Start RabbitMQ container"
        echo "  stop    - Stop RabbitMQ container"
        echo "  restart - Restart RabbitMQ container"
        echo "  logs    - Show RabbitMQ logs"
        echo "  status  - Show container status"
        exit 1
        ;;
esac 