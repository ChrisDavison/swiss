pub fn run(kg: f32) -> Result<(), Box<dyn std::error::Error>> {
    let factor: f32 = 1.78;
    println!("BMI {:.0} is {:.2} Kg", kg, kg * factor.powf(2.0));
    Ok(())
}
