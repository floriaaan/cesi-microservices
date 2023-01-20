use todo::todo_server::{Todo, TodoServer};
use todo::{
    CreateTodoRequest, DeleteTodoRequest, GetTodoRequest, TodoList, TodoReponse, UpdateTodoRequest, DeleteTodoResponse,
};
use tonic::{transport::Server, Request, Response, Status};

pub mod todo {
    tonic::include_proto!("todo");
}

#[derive(Default)]
pub struct TodoService {
    todos: Vec<TodoReponse>,
}

#[tonic::async_trait]
impl Todo for TodoService {
    async fn get_todo(
        &self,
        request: Request<GetTodoRequest>,
    ) -> Result<Response<TodoReponse>, Status> {
        let payload = request.into_inner();

        let message = TodoReponse {
            id: payload.id,
            title: String::from("test"),
            description: String::from("test"),
            done: false,
        };

        Ok(Response::new(message))
    }

    async fn list_todos(&self, _: Request<()>) -> Result<Response<TodoList>, Status> {
        let message = TodoList {
            todos: self.todos.to_vec(),
        };

        Ok(Response::new(message))
    }

    async fn create_todo(
        &self,
        request: Request<CreateTodoRequest>,
    ) -> Result<Response<TodoReponse>, Status> {
        let payload = request.into_inner();

        let todo_item = CreateTodoRequest {
            title: payload.title,
            description: payload.description,
        };

        let message = TodoReponse {
            id: String::from("1"),
            title: todo_item.title,
            description: todo_item.description,
            done: false,
        };

        let todos = &self.todos;
        let mut todos = todos.to_vec();
        todos.push(message.clone());
        

        Ok(Response::new(message))
    }

    async fn update_todo(
        &self,
        request: Request<UpdateTodoRequest>,
    ) -> Result<Response<TodoReponse>, Status> {
        let payload = request.into_inner();

        let message = TodoReponse {
            id: payload.id,
            title: payload.title,
            description: payload.description,
            done: payload.done,
        };

        Ok(Response::new(message))
    }

    async fn delete_todo(&self, _: Request<DeleteTodoRequest>) -> Result<Response<DeleteTodoResponse>, Status> {
        let message = DeleteTodoResponse {
            ok: true,
        };

        Ok(Response::new(message))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "0.0.0.0:50051".parse().unwrap();
    let todo_service = TodoService::default();

    Server::builder()
        .add_service(TodoServer::new(todo_service))
        .serve(addr)
        .await?;

    Ok(())
}
