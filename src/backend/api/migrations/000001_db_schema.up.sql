-- Create the schema
CREATE TABLE categories (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL
);

CREATE TABLE subcategories (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  category_id INTEGER NOT NULL REFERENCES categories(id)
);

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL,
  password VARCHAR(255) NOT NULL,
  company VARCHAR(255) NOT NULL,
  instagram VARCHAR(255) NOT NULL,
  linkedin VARCHAR(255) NOT NULL,
  photo TEXT,
  description TEXT
);

DO $$ BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'project_status_enum') THEN
    CREATE TYPE project_status_enum AS ENUM ('Planejamento', 'Desenvolvimento', 'Finalizado');
  END IF;
END $$;

CREATE TABLE projects (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  description TEXT,
  status project_status_enum NOT NULL,
  user_id INTEGER NOT NULL REFERENCES users(id),
  subcategory_id INTEGER NOT NULL REFERENCES subcategories(id),
  category_id INTEGER NOT NULL REFERENCES categories(id), 
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  photo TEXT,
  local TEXT
);

DO $$ BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'synergy_status_enum') THEN
    CREATE TYPE synergy_status_enum AS ENUM ('Solicitado', 'Em andamento', 'Finalizado', 'Rejeitado');
  END IF;
END $$;

DO $$ BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'synergy_type_enum') THEN
    CREATE TYPE synergy_type_enum AS ENUM ('Aprendizagem', 'Integração', 'Unificação');
  END IF;
END $$;

CREATE TABLE synergies (
  id SERIAL PRIMARY KEY,
  source_project_id INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
  target_project_id INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
  status synergy_status_enum NOT NULL,
  type synergy_type_enum NOT NULL,
  description TEXT
);

DO $$ BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'rating_enum') THEN
    CREATE TYPE rating_enum AS ENUM ('1', '2', '3', '4');
  END IF;
END $$;

CREATE TABLE ratings (
  id SERIAL PRIMARY KEY,
  date DATE DEFAULT CURRENT_DATE,
  level rating_enum NOT NULL,
  user_id INTEGER NOT NULL REFERENCES users(id),
  project_id INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE
);

CREATE TABLE updates (
  id SERIAL PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  description TEXT NOT NULL,
  date DATE NOT NULL,
  created_at DATE,
  synergy_id INTEGER NOT NULL REFERENCES synergies(id) ON DELETE CASCADE
);

DO $$ BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'notification_type_enum') THEN
    CREATE TYPE notification_type_enum AS ENUM ('Solicitação', 'Atualização', 'Novo Projeto', 'Outro');
  END IF;
END $$;

CREATE TABLE notifications (
  id SERIAL PRIMARY KEY,
  received_user_id INTEGER NOT NULL REFERENCES users(id), -- ID do usuário que recebeu a notificação
  sent_user_id INTEGER NOT NULL REFERENCES users(id), -- ID do usuário que enviou a notificação
  received_project_id INTEGER REFERENCES projects(id) ON DELETE CASCADE, -- ID do projeto do usuário que recebeu a notificação
  sent_project_id INTEGER REFERENCES projects(id) ON DELETE CASCADE, -- ID do projeto do usuário que enviou a notificação
  synergy_id INTEGER REFERENCES synergies(id) ON DELETE CASCADE, -- Referência opcional para uma sinergia
  synergy_type synergy_type_enum, -- Tipo de sinergia que a notificação está enviando
  type notification_type_enum NOT NULL,
  title VARCHAR(100) NOT NULL,
  message TEXT NOT NULL,
  status BOOLEAN DEFAULT FALSE, -- FALSE para não lida, TRUE para lida
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

