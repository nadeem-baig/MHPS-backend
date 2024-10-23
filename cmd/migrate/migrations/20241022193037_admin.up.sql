-- Create admins table
CREATE TABLE IF NOT EXISTS admins (
  id SERIAL PRIMARY KEY,
  firstName VARCHAR(255) NOT NULL,
  lastName VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS members (
    aadhaar_number  VARCHAR(12) PRIMARY KEY,  -- Assuming Aadhaar number is unique and 12 characters long
    address         TEXT NOT NULL,
    blood_group     VARCHAR(3) NOT NULL,      -- Blood group is typically small (e.g., "A+")
    contact_number  VARCHAR(15) NOT NULL,     -- To support international phone numbers
    date_of_birth   DATE NOT NULL,            -- Use DATE for storing just the date
    education       VARCHAR(50) NOT NULL,
    email           VARCHAR(255) NOT NULL UNIQUE, -- Assuming email should be unique
    father_name     VARCHAR(255) NOT NULL,
    marital_status  VARCHAR(50) NOT NULL,
    name            VARCHAR(255) NOT NULL,
    std_pin         VARCHAR(6) NOT NULL       -- Postal PIN, assuming it's a 6-digit code in India
);
