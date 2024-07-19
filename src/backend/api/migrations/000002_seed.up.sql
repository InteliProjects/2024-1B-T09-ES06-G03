-- Inserir categorias
INSERT INTO categories (name) VALUES ('Tecnologia'), ('Educação'), ('Saúde'), ('Finanças');

-- Inserir subcategorias
INSERT INTO subcategories (name, category_id) VALUES 
('Desenvolvimento de Software', 1), 
('Redes', 1),
('Ensino a Distância', 2),
('Educação Infantil', 2),
('Medicina', 3),
('Enfermagem', 3),
('Investimentos', 4),
('Contabilidade', 4);

-- Inserir usuários
INSERT INTO users (name, email, password, company, instagram, linkedin) VALUES 
('Alice Silva', 'alice@example.com', 'password123', 'TechCorp', 'https://instagram.com/alice', 'https://linkedin.com/in/alice'),
('Bob Oliveira', 'bob@example.com', 'password123', 'EduTech', 'https://instagram.com/bob', 'https://linkedin.com/in/bob'),
('Carlos Souza', 'carlos@example.com', 'password123', 'HealthCare Inc.', 'https://instagram.com/carlos', 'https://linkedin.com/in/carlos');

-- Inserir projetos
INSERT INTO projects (name, description, status, user_id, subcategory_id) VALUES 
('Sistema de Gestão Escolar', 'Um sistema completo para gestão escolar', 'Em progresso', 1, 3),
('Aplicativo de Telemedicina', 'Aplicativo para consultas médicas online', 'Completo', 3, 5),
('Plataforma de Investimentos', 'Plataforma para gestão de investimentos', 'Em progresso', 2, 7);

-- Inserir sinergias
INSERT INTO synergies (source_project_id, target_project_id, status, description) VALUES 
(1, 2, 'Solicitado', 'Integração de funcionalidades educacionais com telemedicina'),
(2, 3, 'Aprendizagem', 'Aprendendo sobre integração com plataformas financeiras');

-- Inserir avaliações
INSERT INTO ratings (level, user_id, project_id) VALUES 
('4', 1, 1),
('3', 2, 2),
('2', 3, 3);

-- Inserir atualizações
INSERT INTO updates (title, description, date, project_id) VALUES 
('Primeira versão lançada', 'Lançamos a primeira versão do sistema', '2024-05-01', 1),
('Funcionalidade de vídeo chamada adicionada', 'Adicionamos suporte a videochamadas', '2024-04-15', 2),
('Nova parceria com corretora', 'Fechamos parceria com corretora de investimentos', '2024-03-20', 3);

