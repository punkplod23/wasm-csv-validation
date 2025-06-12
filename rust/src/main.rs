use std::error::Error;
use std::fs::File;
use csv;
use std::env;
use serde_json;

use serde::Serialize;
use std::io::{self, Write};

#[derive(Serialize)]
struct Report {
    rows: Vec<RowReport>,
}

#[derive(Serialize)]
struct RowReport {
    row_id: i32,
    errors: Vec<String>,
}


fn read_csv(file_path: &str) -> Result<(), Box<dyn Error>> {
    let file = File::open(file_path)?;
    let mut rdr = csv::Reader::from_reader(file);

    let mut report = Report { rows: Vec::new() };

    // Get the headers
    let headers = rdr.headers()?.clone();

    // Iterate over each record (row) with its index
    for (row_idx, result) in rdr.records().enumerate() {
        let record = result?; // Get the StringRecord
        let mut errors = Vec::new();

        // Iterate over each field in the record
        for (i, field) in record.iter().enumerate() {
            if let Some(header) = headers.get(i) {
                if field.trim().is_empty() {
                    // Use row_idx + 2 to represent the actual row number (1-based, including header)
                    let row_id = (row_idx + 2) as i32;
                    let error_message = format!("Row {}: Empty field in column '{}'", row_id, header);
                    errors.push(error_message);
                }
            } else {
                println!("No header for index {} in record: {:?}", i, record);
            }
        }

        if !errors.is_empty() {
            let row_report = RowReport {
                row_id: (row_idx + 2) as i32,
                errors,
            };
            report.rows.push(row_report);
        }
    }

    let stdout = io::stdout();
    let mut handle = stdout.lock();
    let json_output = serde_json::to_string(&report)?;
    writeln!(handle, "{}", json_output)?;

    Ok(())
}

fn main() {

    // Get the CSV file path from command line arguments
    let args: Vec<String> = env::args().collect();
    if args.len() != 2 {
        eprintln!("Usage: {} <csv_file_path>", args[0]);
        std::process::exit(1);
    }

    let csv_file_path = &args[1];

    if let Err(e) = read_csv(csv_file_path) {
        eprintln!("Error reading CSV: {}", e);
    }
}

