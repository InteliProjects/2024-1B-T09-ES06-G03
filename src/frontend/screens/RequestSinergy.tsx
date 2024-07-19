import React, { useEffect, useState } from "react";
import { Text, View, TouchableHighlight, TouchableWithoutFeedback, Keyboard } from "react-native";
import CardSinergyProjects from "../components/CardSinergyProjects";
import GradientButton from "../components/GradientButton";
import InputComponent from "../components/Input";
import Tabs from "../components/Tabs";
import Dropdown from "../components/Dropdown";
import { LinearGradient } from 'expo-linear-gradient';
import { styled } from 'nativewind';
import { ceoApi, projectApi } from "../services/api";
import { useRoute, useNavigation } from '@react-navigation/native';
import { RouteProp } from '@react-navigation/native';
import { RootStackParamList } from './routes';
import Loading from "../components/Loading";

type EditProfileProps = {
    navigation: any;
};

type Project = {
    id: number;
    name: string;
    user_id?: string;
};

type SynergyRequestPayload = {
    message: string;
    received_project_id: number;
    sent_project_id: number;
    synergy_type: string;
    type: string;
    title: string;
    sent_user_id: number;
    received_user_id: number;
    status: boolean;
    created_at: string;
};

export default function EditProfile({ navigation }: EditProfileProps) {
    const route = useRoute<RouteProp<RootStackParamList, 'ProjectStack'>>();
    const { id } = route.params as unknown as { id: number };

    const StyledLinearGradient = styled(LinearGradient);

    const [loading, setLoading] = useState(true);
    const [step, setStep] = useState(0);
    const [selectedProject, setSelectedProject] = useState<number>();
    const [selectedInterest, setSelectedInterest] = useState<string>("");
    const [description, setDescription] = useState<string>("");
    const [error, setError] = useState<string | null>(null);
    const [project, setProject] = useState<Project>({ id: 0, name: '' });
    const [userProject, setUserProject] = useState<Array<Project>>([]);
    const [userId, setUserId] = useState<number | null>(0);

    const handleNext = () => {
        if (step < 2) {
            setStep(step + 1);
        } else {
            submitSynergyRequest();
            navigation.goBack();
        }
    };

    const getGradientColors = (interest: string) => {
        return selectedInterest === interest ? ['#14514F', '#3A8A88'] : ['white', 'white'];
    };

    const submitSynergyRequest = async () => {
        const payload: SynergyRequestPayload = {
            message: description,
            received_project_id: id,
            sent_project_id: selectedProject,
            synergy_type: selectedInterest,
            type: 'Solicitação',
            title: 'Nova Solicitação de Integração',
            sent_user_id: userId,
            received_user_id: parseInt(project.user_id),
            status: false,
        };

        try {
            const response = await ceoApi.post('/notifications', payload);
            console.log('Notificação solicitada com sucesso:', response.data);
        } catch (err) {
            console.error('Erro ao criar notificação:', err);
            setError('Erro ao criar notificação. Tente novamente.');
        }
    };

    useEffect(() => {
        const fetchProjects = async () => {
            try {
                const response = await projectApi(`/projects/${id}`);
                setProject(response.data);
            } catch (err) {
                console.error('Error fetching data:', err);
                setError('Erro ao buscar dados do projeto.');
            }
        };
    
        const fetchUserProjects = async () => {
            try {
                const response = await projectApi(`/projects/me`);
                setUserProject(response.data);
                setUserId(response.data[0].user_id)
            } catch (err) {
                console.error('Error fetching data:', err);
                setError('Erro ao buscar projetos do usuário.');
            }
        };
    
        const fetchData = async () => {
            setLoading(true);
            await Promise.all([fetchProjects(), fetchUserProjects()]); 
            setLoading(false); 
        };
        fetchData();
    }, [id]);
    

    if (loading) {
        return (
            <View className="w-[100%] h-[100%] justify-center items-center">
                <Loading />
            </View>
        )
    } else {


        return (
            <TouchableWithoutFeedback onPress={Keyboard.dismiss}>
                <View className="h-full w-full bg-white">
                    <View className="flex flex-col h-full justify-between mx-5">
                        <View>
                            <View className="mb-2 items-center">
                                <Text className="font-medium self-start mt-5 mb-5 text-[16px]">Você está solicitando uma sinergia com:</Text>
                                <CardSinergyProjects project={project.name} name={project.user_id || ''} theme={"Empresa"} category={"conservation"} />
                            </View>
                            <Tabs step={step} setStep={setStep} page={'request'} />

                            {step === 0 && (
                                <View className="mb-5">
                                    <Text className="text-[16px]">Qual projeto você quer envolver na sinergia?</Text>
                                    <Dropdown
                                        type=""
                                        label=""
                                        options={userProject.map((project) => ({ label: project.name, value: project.id }))}
                                        selectedValue={selectedProject}
                                        onValueChange={setSelectedProject}
                                        placeholder="Selecione um projeto"
                                    />
                                </View>
                            )}

                            {step === 1 && (
                                <View>
                                    <Text className="text-[16px] mb-10">Qual é o principal interesse com a sinergia?</Text>
                                    <TouchableHighlight onPress={() => setSelectedInterest('Aprendizagem')} className='rounded-2xl w-full my-5'>
                                        <StyledLinearGradient colors={getGradientColors('Aprendizagem')} start={{ x: 0, y: 0 }} end={{ x: 1, y: 0 }} className={`rounded-2xl ${selectedInterest === 'Aprendizagem' ? 'p-[13px]' : 'border-2 border-green-10 p-3'}`}>
                                            <Text className={`text-[18px] font-medium text-center ${selectedInterest === 'Aprendizagem' ? 'text-white' : 'text-green-10'}`}>Aprender</Text>
                                        </StyledLinearGradient>
                                    </TouchableHighlight>

                                    <TouchableHighlight onPress={() => setSelectedInterest('Integração')} className='rounded-2xl w-full my-5'>
                                        <StyledLinearGradient colors={getGradientColors('Integração')} start={{ x: 0, y: 0 }} end={{ x: 1, y: 0 }} className={`rounded-2xl ${selectedInterest === 'Integração' ? 'p-[13px]' : 'border-2 border-green-10 p-3'}`}>
                                            <Text className={`text-[18px] font-medium text-center ${selectedInterest === 'Integração' ? 'text-white' : 'text-green-10'}`}>Integrar</Text>
                                        </StyledLinearGradient>
                                    </TouchableHighlight>

                                    <TouchableHighlight onPress={() => setSelectedInterest('Unificação')} className='rounded-2xl w-full my-5'>
                                        <StyledLinearGradient colors={getGradientColors('Unificação')} start={{ x: 0, y: 0 }} end={{ x: 1, y: 0 }} className={`rounded-2xl ${selectedInterest === 'Unificação' ? 'p-[13px]' : 'border-2 border-green-10 p-3'}`}>
                                            <Text className={`text-[18px] font-medium text-center ${selectedInterest === 'Unificação' ? 'text-white' : 'text-green-10'}`}>Unificar</Text>
                                        </StyledLinearGradient>
                                    </TouchableHighlight>
                                </View>
                            )}

                            {step === 2 && (
                                <View>
                                    <Text className="text-[16px] mb-10">Adicione um comentário para complementar sua solicitação:</Text>
                                    <InputComponent
                                        type="textarea"
                                        label=''
                                        placeholder={"Digite a descrição..."}
                                        maxLength={300}
                                        onChangeText={setDescription}
                                        value={description}
                                    />
                                </View>
                            )}
                        </View>

                        <View>
                            <View className="flex flex-row mb-5">
                                <GradientButton onPress={handleNext} title={step !== 2 ? 'Continuar' : 'Solicitar sinergia'} />
                            </View>
                        </View>
                    </View>
                </View>
            </TouchableWithoutFeedback>
        );
    }
}
