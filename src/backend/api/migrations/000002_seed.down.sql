-- Deletar atualizações
DELETE FROM updates WHERE title IN ('Primeira versão lançada', 'Funcionalidade de vídeo chamada adicionada', 'Nova parceria com corretora');

-- Deletar avaliações
DELETE FROM ratings WHERE level IN ('4', '3', '2') AND user_id IN (1, 2, 3) AND project_id IN (1, 2, 3);

-- Deletar sinergias
DELETE FROM synergies WHERE source_project_id IN (1, 2) AND target_project_id IN (2, 3) AND status IN ('Solicitado', 'Aprendizagem');

-- Deletar projetos
DELETE FROM projects WHERE name IN ('Sistema de Gestão Escolar', 'Aplicativo de Telemedicina', 'Plataforma de Investimentos');

-- Deletar usuários
DELETE FROM users WHERE email IN ('alice@example.com', 'bob@example.com', 'carlos@example.com');

-- Deletar subcategorias
DELETE FROM subcategories WHERE name IN ('Desenvolvimento de Software', 'Redes', 'Ensino a Distância', 'Educação Infantil', 'Medicina', 'Enfermagem', 'Investimentos', 'Contabilidade');

-- Deletar categorias
DELETE FROM categories WHERE name IN ('Tecnologia', 'Educação', 'Saúde', 'Finanças');

