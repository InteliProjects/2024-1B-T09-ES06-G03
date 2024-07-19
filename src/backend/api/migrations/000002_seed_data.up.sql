-- Insert data into categories
INSERT INTO categories (name) VALUES
('conservation'),
('productivity'),
('health'),
('diversity'),
('environmentalImpact'),
('integrity');

-- Insert data into subcategories
INSERT INTO subcategories (name, category_id) VALUES
('Cultivo e Regeneração', 1),
('Capacitar para Ampliar Acesso', 2),
('Bem-Estar e Saúde Mental', 3),
('Raça', 4),
('Mulheres', 4),
('DE&I Geral', 4),
('Descarbonização', 5),
('Economia Circular', 5);

-- Insert data into users
INSERT INTO users (name, email, password, company, instagram, linkedin, photo, description) VALUES
('João Silva', 'joao@example.com', 'hashed_password', 'Tech Corp', '@joaosilva', 'linkedin.com/in/joaosilva', 'https://media.licdn.com/dms/image/D4E03AQGJYYarlpwQbg/profile-displayphoto-shrink_400_400/0/1678716774062?e=1723680000&v=beta&t=a5OVAKEMfCZhUPIDmMwk2Jsi80tZttCn0uygmtCMHig', 'Projeto com o intuito de um mundo melhor'),
('Maria Oliveira', 'maria@example.com', 'hashed_password', 'Health Inc', '@mariaoliveira', 'linkedin.com/in/mariaoliveira', 'https://media.licdn.com/dms/image/D4E03AQGJYYarlpwQbg/profile-displayphoto-shrink_400_400/0/1678716774062?e=1723680000&v=beta&t=a5OVAKEMfCZhUPIDmMwk2Jsi80tZttCn0uygmtCMHig', 'Projeto com o intuito de um mundo melhor'),
('Pedro Santos', 'pedro@example.com', 'hashed_password', 'Edu Tech', '@pedrosantos', 'linkedin.com/in/pedrosantos', 'https://media.licdn.com/dms/image/D4E03AQGJYYarlpwQbg/profile-displayphoto-shrink_400_400/0/1678716774062?e=1723680000&v=beta&t=a5OVAKEMfCZhUPIDmMwk2Jsi80tZttCn0uygmtCMHig', 'Projeto com o intuito de um mundo melhor');

-- Insert data into projects
INSERT INTO projects (name, description, status, user_id, category_id, subcategory_id, photo, local) VALUES
('Website Institucional', 'Desenvolvimento de um site institucional para a empresa.', 'Planejamento', 1, 1, 1, 'https://media.licdn.com/dms/image/D4E03AQGJYYarlpwQbg/profile-displayphoto-shrink_400_400/0/1678716774062?e=1723680000&v=beta&t=a5OVAKEMfCZhUPIDmMwk2Jsi80tZttCn0uygmtCMHig', 'Share Butantã'),
('Chatbot Médico', 'Desenvolvimento de um chatbot para consultas médicas.', 'Desenvolvimento', 2, 2, 2, 'https://media.licdn.com/dms/image/D4E03AQGJYYarlpwQbg/profile-displayphoto-shrink_400_400/0/1678716774062?e=1723680000&v=beta&t=a5OVAKEMfCZhUPIDmMwk2Jsi80tZttCn0uygmtCMHig', 'Inteli - Instituto de Tecnologia e Liderança'),
('Plataforma de Ensino', 'Plataforma online para ensino fundamental.', 'Finalizado', 3, 3, 3, 'https://media.licdn.com/dms/image/D4E03AQGJYYarlpwQbg/profile-displayphoto-shrink_400_400/0/1678716774062?e=1723680000&v=beta&t=a5OVAKEMfCZhUPIDmMwk2Jsi80tZttCn0uygmtCMHig', 'Av. Corifeu de Azevedo Marques, 4160 - Jaguaré, São Paulo - SP'),
('Projeto Agro Sustentável', 'Implementação de técnicas agrícolas sustentáveis.', 'Planejamento', 1, 1, 1, 'https://example.com/photo1.jpg', 'Fazenda Verde'),
('Automação de Processos', 'Automatização de processos industriais.', 'Desenvolvimento', 2, 2, 2, 'https://example.com/photo2.jpg', 'Tech Park'),
('Plataforma de Telemedicina', 'Serviço de consultas médicas online.', 'Finalizado', 3, 3, 3, 'https://example.com/photo3.jpg', 'Saúde Online'),
('Programa de Inclusão', 'Iniciativa para inclusão de minorias.', 'Planejamento', 1, 4, 4, 'https://example.com/photo4.jpg', 'Centro Comunitário'),
('Empoderamento Feminino', 'Projetos para o empoderamento de mulheres.', 'Desenvolvimento', 2, 4, 5, 'https://example.com/photo5.jpg', 'Instituto Mulher'),
('Economia Circular', 'Estratégias para uma economia circular.', 'Finalizado', 3, 5, 8, 'https://example.com/photo6.jpg', 'EcoLab'),
('Projeto de Descarbonização', 'Redução de emissões de carbono.', 'Planejamento', 1, 5, 7, 'https://example.com/photo7.jpg', 'Green Energy'),
('Educação para Todos', 'Programa de educação inclusiva.', 'Desenvolvimento', 2, 2, 2, 'https://example.com/photo8.jpg', 'EduCenter'),
('Saúde Mental no Trabalho', 'Iniciativas para melhorar a saúde mental no ambiente de trabalho.', 'Finalizado', 3, 3, 3, 'https://example.com/photo9.jpg', 'WorkHealth'),
('Conservação de Florestas', 'Projeto para a conservação de áreas florestais.', 'Planejamento', 1, 1, 1, 'https://example.com/photo10.jpg', 'Floresta Viva'),
('Acesso à Saúde', 'Programa para ampliar o acesso à saúde.', 'Desenvolvimento', 2, 2, 2, 'https://example.com/photo11.jpg', 'Health4All'),
('Projeto de Diversidade', 'Iniciativa para promover a diversidade no ambiente corporativo.', 'Finalizado', 3, 4, 6, 'https://example.com/photo12.jpg', 'CorpDiversity'),
('Inovação em Agricultura', 'Novas técnicas para a agricultura moderna.', 'Planejamento', 1, 1, 1, 'https://example.com/photo13.jpg', 'AgroTech'),
('Economia Sustentável', 'Projetos de economia sustentável.', 'Desenvolvimento', 2, 5, 8, 'https://example.com/photo14.jpg', 'SustEco'),
('Inclusão Digital', 'Programa de inclusão digital para comunidades carentes.', 'Finalizado', 3, 2, 2, 'https://example.com/photo15.jpg', 'DigitalInclusion'),
('Bem-Estar Animal', 'Iniciativas para o bem-estar animal.', 'Planejamento', 1, 3, 3, 'https://example.com/photo16.jpg', 'AnimalCare'),
('Projeto Verde', 'Iniciativa para aumentar áreas verdes urbanas.', 'Desenvolvimento', 2, 5, 7, 'https://example.com/photo17.jpg', 'GreenCity'),
('Empoderamento Juvenil', 'Projetos para empoderar jovens.', 'Finalizado', 3, 4, 5, 'https://example.com/photo18.jpg', 'YouthPower'),
('Agricultura Urbana', 'Projetos de agricultura urbana.', 'Planejamento', 1, 1, 1, 'https://example.com/photo19.jpg', 'UrbanFarm'),
('Melhoria de Produtividade', 'Iniciativas para aumentar a produtividade.', 'Desenvolvimento', 2, 2, 2, 'https://example.com/photo20.jpg', 'ProductivityBoost'),
('Saúde e Bem-Estar', 'Programas de saúde e bem-estar.', 'Finalizado', 3, 3, 3, 'https://example.com/photo21.jpg', 'WellBeing'),
('Projeto de Integração', 'Iniciativas para integração de sistemas.', 'Planejamento', 1, 2, 2, 'https://example.com/photo22.jpg', 'IntegrationLab'),
('Conservação de Recursos', 'Projetos para conservação de recursos naturais.', 'Desenvolvimento', 2, 1, 1, 'https://example.com/photo23.jpg', 'ResourceConservation'),
('Promoção de Diversidade', 'Programas para promoção da diversidade.', 'Finalizado', 3, 4, 6, 'https://example.com/photo24.jpg', 'DiversityHub'),
('Desenvolvimento Sustentável', 'Projetos de desenvolvimento sustentável.', 'Planejamento', 1, 5, 8, 'https://example.com/photo25.jpg', 'SustainableDev');


-- Insert data into synergies
INSERT INTO synergies (source_project_id, target_project_id, status, type, description) VALUES
(1, 2, 'Em andamento', 'Integração', 'Integração entre o site institucional e o chatbot médico.'),
(2, 3, 'Finalizado', 'Unificação', 'Unificação da plataforma de ensino com o chatbot médico.');


-- Insert data into ratings
-- ratings_level need to be on (-2, 2)* scale
INSERT INTO ratings (date, level, user_id, project_id) VALUES
(CURRENT_DATE, '2', 1, 1),
(CURRENT_DATE, '2', 2, 2),
(CURRENT_DATE, '2', 3, 3),
(CURRENT_DATE, '2', 1, 4),
(CURRENT_DATE, '2', 2, 5),
(CURRENT_DATE, '2', 3, 6),
(CURRENT_DATE, '2', 1, 7),
(CURRENT_DATE, '2', 2, 8),
(CURRENT_DATE, '2', 3, 9),
(CURRENT_DATE, '2', 1, 10),
(CURRENT_DATE, '2', 2, 11),
(CURRENT_DATE, '2', 3, 12),
(CURRENT_DATE, '2', 1, 13),
(CURRENT_DATE, '2', 2, 14),
(CURRENT_DATE, '2', 3, 15),
(CURRENT_DATE, '2', 1, 16),
(CURRENT_DATE, '2', 2, 17),
(CURRENT_DATE, '2', 3, 18),
(CURRENT_DATE, '2', 1, 19),
(CURRENT_DATE, '2', 2, 20),
(CURRENT_DATE, '2', 3, 21),
(CURRENT_DATE, '2', 1, 22),
(CURRENT_DATE, '2', 2, 23),
(CURRENT_DATE, '2', 3, 24),
(CURRENT_DATE, '2', 1, 25),
(CURRENT_DATE, '2', 2, 26),
(CURRENT_DATE, '2', 3, 27),
(CURRENT_DATE, '2', 1, 28);

-- Insert data into updates
INSERT INTO updates (title, description, date, created_at, synergy_id) VALUES
('Lançamento do Site', 'Lançamento do site institucional.', CURRENT_DATE, CURRENT_DATE, 1),
('Atualização do Chatbot', 'Atualização com novas funcionalidades.', CURRENT_DATE, CURRENT_DATE, 2);

-- Insert data into notifications
INSERT INTO notifications (
  received_user_id, 
  sent_user_id, 
  received_project_id, 
  sent_project_id, 
  synergy_type, 
  type, 
  title, 
  message, 
  status, 
  created_at
) VALUES
(1, 2, 1, 2, 'Integração', 'Solicitação', 'Nova Solicitação de Integração', 'Solicitação de integração de projeto recebida.', FALSE, CURRENT_TIMESTAMP),
(2, 3, 2, 3, 'Unificação', 'Atualização', 'Unificação de Projetos', 'Seu projeto foi unificado com outro projeto.', FALSE, CURRENT_TIMESTAMP),
(3, 1, 3, 1, 'Aprendizagem', 'Novo Projeto', 'Novo Projeto Adicionado', 'Um novo projeto de ensino foi criado e pode ser do seu interesse.', FALSE, CURRENT_TIMESTAMP);


