use chrono::prelude::*;
use chrono::Days;

pub fn start() -> Result<(), Box<dyn std::error::Error>> {
    let now = Utc::now().date_naive();
    if let Some(ok) = now.checked_sub_days(Days::new(now.weekday() as u64)) {
        println!("{}", ok);
    }
    Ok(())
}

pub fn start_and_end() -> Result<(), Box<dyn std::error::Error>> {
    let now = Utc::now().date_naive();
    let start = now.checked_sub_days(Days::new(now.weekday() as u64));
    let end = start.and_then(|x| x.checked_add_days(Days::new(6)));
    if let (Some(start), Some(end)) = (start, end) {
        println!("{} to {}", start, end);
    }
    Ok(())
}
