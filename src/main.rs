use rocket::serde::json::Json;

#[macro_use]
extern crate rocket;

#[derive(FromForm, Default)]
pub struct SumParams {
    number1: Option<i32>,
    number2: Option<i32>,
}

#[get("/sum?<params..>")]
fn index(params: SumParams) -> Json<i32> {
    let number1 = params.number1.unwrap();
    let number2 = params.number2.unwrap();
    return Json(number1 + number2);
}

#[launch]
fn rocket() -> _ {
    rocket::build().mount("/", routes![index])
}
