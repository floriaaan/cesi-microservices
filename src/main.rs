use rocket::{ serde::json::Json};
use serde::Deserialize;

#[macro_use]
extern crate rocket;

#[derive(FromForm, Default, Deserialize)]
pub struct SumParams {
    number1: Option<i32>,
    number2: Option<i32>,
}

#[get("/sum?<params..>")]
fn get(params: SumParams) -> Json<i32> {
    let number1 = params.number1.unwrap();
    let number2 = params.number2.unwrap();
    return Json(number1 + number2);
}

#[post("/sum", data = "<params>")]
fn post(params: Json<SumParams>) -> Json<i32> {
    let number1 = params.number1.unwrap();
    let number2 = params.number2.unwrap();
    return Json(number1 + number2);
}

#[launch]
fn rocket() -> _ {
    rocket::build().mount("/", routes![get, post])
}
