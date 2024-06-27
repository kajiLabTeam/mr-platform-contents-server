\c mr-platform-contents;

-- Enum Type for Contents
CREATE TYPE content_type AS ENUM (
    'html2d', 
    'model3d'
);

-- Enum Type for Text Type
CREATE TYPE text_type AS ENUM (
    'html',
    'markdown'
);

-- Ceraate PublicSpaces Table
CREATE TABLE public_spaces (
    id VARCHAR(36) PRIMARY KEY,
    organization_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create Contents Table
CREATE TABLE contents (
    id VARCHAR(36) PRIMARY KEY,
    public_space_id VARCHAR(36) REFERENCES public_spaces(id),
    type content_type NOT NULL,
    domain VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create html2d Table
CREATE TABLE html2d (
    id VARCHAR(36) PRIMARY KEY,
    content_id VARCHAR(36) REFERENCES contents(id),
    location_x DOUBLE PRECISION NOT NULL,
    location_y DOUBLE PRECISION NOT NULL,
    location_z DOUBLE PRECISION NOT NULL,
    rotation_row DOUBLE PRECISION NOT NULL,
    rotation_pitch DOUBLE PRECISION NOT NULL,
    rotation_yaw DOUBLE PRECISION NOT NULL,
    text_type text_type NOT NULL,
    text_url VARCHAR(255) NOT NULL,
    style_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create model3d Table
CREATE TABLE model3d (
    id VARCHAR(36) PRIMARY KEY,
    content_id VARCHAR(36) REFERENCES contents(id),
    location_x DOUBLE PRECISION NOT NULL,
    location_y DOUBLE PRECISION NOT NULL,
    location_z DOUBLE PRECISION NOT NULL,
    rotation_row DOUBLE PRECISION NOT NULL,
    rotation_pitch DOUBLE PRECISION NOT NULL,
    rotation_yaw DOUBLE PRECISION NOT NULL,
    file_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);