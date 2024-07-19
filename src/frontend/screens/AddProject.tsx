import React, { useState } from "react";
import { Text, View, StyleSheet, TouchableOpacity, Image, Dimensions, Alert } from "react-native";
import * as ImagePicker from 'expo-image-picker';
import InputComponent from '../components/Input';
import Dropdown from '../components/Dropdown';
import { KeyboardAwareScrollView } from 'react-native-keyboard-aware-scroll-view';
import GradientButton from "../components/GradientButton";
import { projectApi } from "../services/api";

const { width } = Dimensions.get('window');

export default function AddProject({ navigation }) {
    const [selectedThemeValue, setSelectedThemeValue] = useState(null);
    const [selectedSubthemeValue, setSelectedSubthemeValue] = useState(null);
    const [projectName, setProjectName] = useState('');
    const [description, setDescription] = useState('');
    const [photo, setPhoto] = useState(null);
    const [local, setLocal] = useState('');

    const categoryOptions = [
        { label: 'Conservação do Planeta', value: 1 },
        { label: 'Produtividade e Competitividade', value: 2 },
        { label: 'Bem estar, saúde e felicidade', value: 3 },
        { label: 'Dignidade e Integridade', value: 4 },
        { label: 'Redução do impacto Ambiental', value: 5 },
        { label: 'Integridade e práticas éticas', value: 6 },
    ];
    
    const subcategoryOptions = {
        1: [
            { label: 'Cultivo e Regeneração', value: 1 },
        ],
        2: [
            { label: 'Capacitar para Ampliar Acesso', value: 2 },
        ],
        3: [
            { label: 'Bem-Estar e Saúde Mental', value: 3 },
        ],
        4: [
            { label: 'Raça', value: 4 },
            { label: 'Mulheres', value: 5 },
            { label: 'DE&I Geral', value: 6 },
        ],
        5: [
            { label: 'Descarbonização', value: 7 },
            { label: 'Economia Circular', value: 8 },
        ],
        6: []
    };

    const handleCreateProject = async () => {
        const newProject = {
            name: projectName,
            description: description,
            user_id: 7, 
            status: 'Planejamento', // Use um valor válido conforme o enum
            subcategory_id: Number(selectedSubthemeValue), // Certifique-se de que é um númerostatus: 'published', // Use um valor válido conforme o enum
            category_id: Number(selectedThemeValue),
            photo: photo || '', // Ajuste conforme necessário
            local: local,
            
            // Ajuste conforme necessário
        };

        console.log('Creating project with data:', newProject);

        try {
            const response = await projectApi.post('/projects', newProject);
            if (response.status === 201) {
                console.log('Project created successfully');
                Alert.alert('Sucesso', 'Projeto criado com sucesso!');
                navigation.goBack(); // Retorna para a tela anterior após a criação do projeto
            } else {
                console.log('Unexpected response:', response);
                Alert.alert('Erro', `Erro ao criar projeto. Código: ${response.status}`);
            }
        } catch (err) {
            if (err.response) {
                // Erros que possuem uma resposta do servidor
                console.error('Error creating project:', err.response.data);
                Alert.alert('Erro', `Erro ao criar projeto: ${err.response.data.message}`);
            } else if (err.request) {
                // Erros que ocorreram durante a requisição, mas não houve resposta
                console.error('Error creating project:', err.request);
                Alert.alert('Erro', 'Erro ao criar projeto: Nenhuma resposta do servidor.');
            } else {
                // Outros erros
                console.error('Error creating project:', err.message);
                Alert.alert('Erro', `Erro ao criar projeto: ${err.message}`);
            }
        }
    };

    const pickImage = async () => {
        console.log('Requesting media library permissions...');
        const { status } = await ImagePicker.requestMediaLibraryPermissionsAsync();
        if (status !== 'granted') {
            alert('Desculpe, precisamos de permissão para acessar a galeria!');
            return;
        }

        console.log('Launching image picker...');
        let result = await ImagePicker.launchImageLibraryAsync({
            mediaTypes: ImagePicker.MediaTypeOptions.Images,
            allowsEditing: true,
            aspect: [4, 3],
            quality: 1,
        });

        if (!result.canceled) {
            console.log('Image picked:', result.assets[0].uri);
            setPhoto(result.assets[0].uri);
        }
    };

    return (
        <KeyboardAwareScrollView
            style={{ flex: 1 }}
            resetScrollToCoords={{ x: 0, y: 0 }}
            contentContainerStyle={{ flexGrow: 1 }}
            scrollEnabled={true}
        >
            <View style={styles.container}>
                <TouchableOpacity style={styles.imageUpload} onPress={pickImage}>
                    {photo ? (
                        <Image source={{ uri: photo }} style={styles.uploadIcon} />
                    ) : (
                        <>
                            <Image source={require('../assets/imageIcon.png')} style={styles.uploadIcon} />
                            <Image source={require('../assets/uploadIcon.png')} style={styles.smallUploadIcon} />
                        </>
                    )}
                </TouchableOpacity>
                <InputComponent
                    label="Nome do projeto"
                    type="input"
                    placeholder="Digite o nome do projeto..."
                    maxLength={50}
                    onChangeText={setProjectName} 
                    value={''}
                />
                <InputComponent
                    label="Localização"
                    type="input"
                    placeholder="Digite a localização do projeto..."
                    maxLength={50}
                    onChangeText={setLocal} 
                    value={''}
                />
                <Dropdown
                    type=""
                    label="Tema"
                    options={categoryOptions}
                    selectedValue={selectedThemeValue}
                    onValueChange={setSelectedThemeValue}
                    placeholder="Selecione um tema"
                />
                <Dropdown
                    type=""
                    label="Subtema"
                    options={subcategoryOptions[selectedThemeValue] || []}
                    selectedValue={selectedSubthemeValue}
                    onValueChange={setSelectedSubthemeValue}
                    placeholder="Selecione um subtema"
                />
                <InputComponent
                    label="Descrição"
                    type="textarea"
                    placeholder="Crie uma descrição para o projeto..."
                    maxLength={2000}
                    onChangeText={setDescription}
                    value={''}
                />
                <View style={styles.buttonContainer}>
                    <GradientButton title="Criar Projeto" onPress={handleCreateProject} />
                </View>
            </View>
        </KeyboardAwareScrollView>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'center',
        padding: 10,
        backgroundColor: '#fff',
    },
    input: {
        width: width * 0.9,
        marginBottom: width * 0.01,
        backgroundColor: '#fff',
        borderRadius: 5,
        zIndex: 1,
    },
    dropdown: {
        width: width * 0.9,
        marginBottom: width * 0.01,
        backgroundColor: '#fff',
        borderRadius: 5,
        zIndex: 1,
    },
    button: {
        width: '75%',
        height: 50,
        borderRadius: 10,
        overflow: 'hidden',
        marginTop: 10,
        backgroundColor: '#08d4c4',
    },
    gradient: {
        flex: 1,
        justifyContent: 'center',
        alignItems:'center',
    },
    buttonText: {
        color: '#fff',
        fontSize: 16,
    },
    imageUpload: {
        width: 100,
        height: 100,
        borderRadius: 50,
        backgroundColor: '#ddd',
        justifyContent: 'center',
        alignItems: 'center',
        marginTop: 20,
    },
    uploadIcon: {
        width: 100,
        height: 100,
        borderRadius: 50,
    },
    smallUploadIcon: {
        width: 30,
        height: 30,
        position: 'absolute',
        right: 0,
        bottom: 0,
    },
    buttonContainer: {
        width: '90%',
        marginTop: 10,
        alignItems: 'center',
        justifyContent: 'center',
    }
});
