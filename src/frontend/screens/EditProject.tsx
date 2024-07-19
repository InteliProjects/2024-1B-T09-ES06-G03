import React, { useEffect, useState } from "react";
import { Text, View, TouchableOpacity, Image } from "react-native";
import InputComponent from "../components/Input";
import Dropdown from "../components/Dropdown";
import { useRoute } from '@react-navigation/native';
import { RouteProp } from '@react-navigation/native';
import { RootStackParamList } from './routes';
import Loading from "../components/Loading";
import { projectApi } from "../services/api";
import GradientButton from "../components/GradientButton";
import { KeyboardAwareScrollView } from 'react-native-keyboard-aware-scroll-view';

type UpdateProjectPayload = {
    name: string;
    description: string;
    category_id: number;
    subcategory_id: number;
    status: string;
    photo: string;
    user_id: number;
};

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

const statusOptions = [
    { label: 'Planejamento', value: 1 },
    { label: 'Desenvolvimento', value: 2 },
    { label: 'Finalizado', value: 3 },
];

export default function EditProject({ navigation }) {
    const route = useRoute<RouteProp<RootStackParamList, 'ProjectStack'>>();
    const { id } = route.params as unknown as { id: number };

    const [selectedTheme, setSelectedTheme] = useState<number | null>(null);
    const [projectName, setProjectName] = useState('');
    const [description, setDescription] = useState('');
    const [loading, setLoading] = useState(true);
    const [selectedSubtheme, setSelectedSubtheme] = useState<number | null>(null);
    const [currentStatus, setCurrentStatus] = useState<number | null>(null);
    const [userId, setUserId] = useState<number | null>(null);

    useEffect(() => {
        const fetchUserProjects = async () => {
            try {
                const response = await projectApi.get('/projects/me');
                setUserId(response.data[0].user_id);
            } catch (err) {
                console.error('Error fetching user projects:', err);
            }
        }

        const fetchProject = async () => {
            try {
                const response = await projectApi(`/projects/${id}`);
                console.log('Project fetched:', response.data);
                setProjectName(response.data.name);
                setDescription(response.data.description);
                setSelectedTheme(response.data.category_id);
                setSelectedSubtheme(response.data.subcategory_id);
                setCurrentStatus(response.data.status === 'Planejamento' ? 1 : response.data.status === 'Desenvolvimento' ? 2 : 3);
            } catch (err) {
                console.error('Error fetching project:', err);
            } finally {
                setLoading(false);
            }
        };

        fetchUserProjects();
        fetchProject();
    }, [id]);

    const updateProject = async () => {
        const payload: UpdateProjectPayload = {
            name: projectName,
            description: description,
            category_id: selectedTheme || 0,
            subcategory_id: selectedSubtheme || 0,
            status: currentStatus === 1 ? 'Planejamento' : currentStatus === 2 ? 'Desenvolvimento' : 'Finalizado',
            photo: 'url',
            user_id: userId,
        };

        try {
            const response = await projectApi.put(`/projects/${id}`, payload);
            console.log('Projeto atualizado com sucesso:', response.data);
            navigation.goBack();
        } catch (err) {
            console.error('Erro ao atualizar projeto:', err);
        }
    };

    if (loading) {
        return <Loading />;
    }

    return (
        <KeyboardAwareScrollView>
            <View style={{ alignItems: 'center' }}>
                <TouchableOpacity style={{ borderRadius: 20, justifyContent: 'center', alignItems: 'center', marginTop: 20, width: 100, height: 100 }}>
                    <Image source={require('../assets/imageIcon.png')} style={{ width: '100%', height: '100%' }} />
                    <Image source={require('../assets/uploadIcon.png')} style={{ width: 30, height: 30, position: 'absolute', bottom: 0, right: 0 }} />
                </TouchableOpacity>

                <View style={{ marginTop: 20, gap: 16 }}>
                    <InputComponent
                        label="Nome do projeto"
                        type="input"
                        value={projectName}
                        maxLength={50}
                        onChangeText={setProjectName}
                        placeholder={projectName}
                    />

                    <Dropdown
                        label="Status"
                        type="green"
                        options={statusOptions}
                        selectedValue={currentStatus}
                        onValueChange={setCurrentStatus}
                        placeholder={statusOptions.find(option => option.value === currentStatus)?.label || "Selecione um status"}
                    />

                    <Dropdown
                        label="Tema"
                        type="green"
                        options={categoryOptions}
                        selectedValue={selectedTheme}
                        onValueChange={(value) => {
                            setSelectedTheme(value);
                            setSelectedSubtheme(null);
                        }}
                        placeholder={categoryOptions.find(option => option.value === selectedTheme)?.label || "Selecione um tema"}
                    />

                    <Dropdown
                        label="Subtema"
                        type="green"
                        options={subcategoryOptions[selectedTheme] || []}
                        selectedValue={selectedSubtheme}
                        onValueChange={setSelectedSubtheme}
                        placeholder={subcategoryOptions[selectedTheme]?.find(option => option.value === selectedSubtheme)?.label || "Selecione um subtema"}
                    />

                    <InputComponent
                        label="Descrição"
                        type="textarea"
                        value={description}
                        maxLength={2000}
                        onChangeText={setDescription}
                        placeholder={description}
                    />
                </View>
                <View style={{ marginTop: 20, marginBottom: 20, width: '90%' }}>
                    <GradientButton title="Salvar alterações" onPress={updateProject} />
                </View>
            </View>
        </KeyboardAwareScrollView>
    )
}
