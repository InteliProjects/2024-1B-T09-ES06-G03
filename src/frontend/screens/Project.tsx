import React, { useEffect, useState } from "react";
import { Text, View, TouchableOpacity, ImageBackground } from "react-native";
import { LinearGradient } from 'expo-linear-gradient';
import CeoAvatar from "../components/CeoAvatar";
import Hands from "../assets/handsWhite.svg";
import Map from "../assets/map.svg";
import Check from "../assets/checkIcon.svg";
import Flag from "../assets/flagIcon.svg";
import { projectApi } from "../services/api";
import { useRoute } from '@react-navigation/native';
import { RouteProp } from '@react-navigation/native';
import { RootStackParamList } from './routes';
import { set } from "react-hook-form";

function filterProjects(projects: any[], id: number) {
  return projects.filter((project) => project.id === id);
}

interface Project {
  id: number,
  name: string,
  description: string,
  status: string,
  user_id: number,
  category_id: number,
  subcategory_id: number,
  created_at: any,
  updated_at: any,
  photo: string,
  local: string
  subcategory_name: string,
  category_name: string,
  synergy_count: number,
  interested_count: number
}

export default function Project({ navigation }) {
  const [project, setProject] = useState<Project>({
    id: 0,
    name: '',
    description: '',
    status: '',
    user_id: 0,
    category_id: 0,
    subcategory_id: 0,
    created_at: '',
    updated_at: '',
    photo: '',
    local: '',
    subcategory_name: '',
    category_name: '',
    synergy_count: 0,
    interested_count: 0
  });
  const [interested, setInterested] = useState({ interested_users: [] });
  const [userId, setUserId] = useState(0);

  const route = useRoute<RouteProp<RootStackParamList, 'ProjectStack'>>();
  const { id } = route.params as unknown as { id: number };

  const goToRequest = () => {
    navigation.navigate('ProjectStack', {
      screen: 'RequestSinergy',
      params: { id: id }
    })
  };

  const goToEdit = () => {
    navigation.navigate('ProjectStack', {
      screen: 'EditProject',
      params: { id: id }
    })
  };

  useEffect(() => {

    const fetchUserProjects = async () => {
      try {
        const response = await projectApi.get(`/projects/me`);
        setUserId(response.data[0].user_id);

      } catch (error) {
        console.log(error);
      }
    };

    const fetchProjects = async () => {
      try {
        const response = await projectApi.get(`/projects`);
        const project = filterProjects(response.data, id);
        return project[0];

      } catch (error) {
        console.log(error);
      }
    };

    const fetchInterested = async () => {
      try {
        const response = await projectApi.get(`/projects/${id}/interested`);
        return response.data;
      } catch (error) {
        console.log(error);
      }
    }

    const fetchData = async () => {
      const projects = await fetchProjects();
      const interestedData = await fetchInterested();
      fetchUserProjects();
      setInterested({ interested_users: interestedData });
      setProject(projects);
    }

    fetchData();
  }, [id]);

  let content;
  if (project.status === 'Planejamento') {
    content = <>
      <View className="flex-row justify-between w-[330px]">
        <Text className="text-green-10">Planejamento</Text>
        <Text className="text-gray-4">Desenvolvimento</Text>
        <Text className="text-gray-4">Finalizado</Text>
      </View>
      <View className="flex-row mt-2 items-center">
        <View className="w-8 h-8 rounded-full bg-green-10 items-center justify-center">
          <Check height={15} width={15} />
        </View>
        <View className="h-1 bg-gray-4 w-28" />
        <View className="w-8 h-8 rounded-full bg-gray-4 items-center justify-center">
        </View>
        <View className="h-1 bg-gray-4 w-28" />
        <View className="w-8 h-8 rounded-full bg-gray-4 items-center justify-center">
        </View>
      </View>
    </>;
  } else if (project.status === 'Desenvolvimento') {
    content = <>
      <View className="flex-row justify-between">
        <Text className="text-green-10">Planejamento</Text>
        <Text className="text-green-10">Desenvolvimento</Text>
        <Text className="text-gray-4 ">Finalizado</Text>
      </View>
      <View className="flex-row mt-2 items-center">
        <View className="w-8 h-8 rounded-full bg-green-10 items-center justify-center">
          <Check height={15} width={15} />
        </View>
        <View className="h-1 bg-green-10 w-28" />
        <View className="w-8 h-8 rounded-full bg-green-10 items-center justify-center">
          <Check height={15} width={15} />
        </View>
        <View className="h-1 bg-gray-4 w-28" />
        <View className="w-8 h-8 rounded-full bg-gray-4 items-center justify-center">
        </View>
      </View>
    </>
  } else if (project.status === 'Finalizado') {
    content = <>
      <View className="flex-row justify-between">
        <Text className="text-green-10">Planejamento</Text>
        <Text className="text-green-10">Desenvolvimento</Text>
        <Text className="text-green-10">Finalizado</Text>
      </View>
      <View className="flex-row mt-2 items-center">
        <View className="w-8 h-8 rounded-full bg-green-10 items-center justify-center">
          <Check height={15} width={15} />
        </View>
        <View className="h-1 bg-green-10 w-28" />
        <View className="w-8 h-8 rounded-full bg-green-10 items-center justify-center">
          <Check height={15} width={15} />
        </View>
        <View className="h-1 bg-green-10 w-28" />
        <View className="w-8 h-8 rounded-full bg-green-10 items-center justify-center">
          <Check height={15} width={15} />
        </View>
      </View>
    </>
  }

  return (
    <View className="flex-1 flex-col ">
      <View className="h-[16%]">
        <ImageBackground className="w-[100%] h-[100%] items-center justify-center bg-green-10 mt-6" />
      </View>
      <View className="w-[100%]">
        <View className="flex-row justify-center mr-4">
          <ImageBackground className="w-32 h-32  bg-black rounded-full" source={project.photo}>
            <View className="items-end justify-end h-[100%]">
              <CeoAvatar size="w-12 h-12" link={project.photo} />
            </View>
          </ImageBackground>
          <View className="items-end justify-end ml-14">

            {userId === project.user_id ? (
              <LinearGradient
                className="w-40 h-9 rounded-2xl"
                colors={['#ABABAB', '#B3B3B3']}
                start={{ x: 0, y: 0 }}
                end={{ x: 1, y: 0 }}
              >
                <TouchableOpacity className="justify-center items-center flex-row w-full h-full" onPress={goToEdit}>
                  <Text className="text-white ml-2 text-lg font-inter">Editar projeto</Text>
                </TouchableOpacity>
              </LinearGradient>
            ) : (
              <LinearGradient
                className="w-40 h-9 rounded-2xl"
                colors={['rgba(20, 81, 79, 1)', 'rgba(58, 138, 136, 1)']}
                start={{ x: 0, y: 0 }}
                end={{ x: 1, y: 0 }}
              >
                <TouchableOpacity className="w-full justify-center items-center h-full flex-row" onPress={goToRequest}>
                  <Hands height={20} width={20} />
                  <Text className="text-white ml-2 text-lg font-inter">Criar Sinergia</Text>
                </TouchableOpacity>
              </LinearGradient>
            )}
          </View>
        </View>
        <View className="flex-row mt-2 w-[100%]  ">
          <View className="w-[40%] ml-6">
            <Text className="font-bold text-xl ">{project.name}</Text>
            <Text className="text-gray-4">{project.subcategory_name}</Text>
            <Text className="text-gray-4">{project.category_name}</Text>
            <View className="flex-row mt-2">
              <Map height={20} width={20} />
              <Text className="text-gray-4 ml-2">{project.local}</Text>
            </View>
          </View>
          <View className="w-[60%] flex-row">
            <Text><Text className="font-bold text-sm">{project.synergy_count}</Text> sinergias</Text>
            <Text className="ml-2"><Text className="font-bold text-sm">{project.interested_count}</Text> interressados</Text>
          </View>
        </View>
        <View className="mt-2 w-[100%] ml-6 justify-center">
          <Text className="text-base">{project.description} </Text>
        </View>
        <View className="mt-8 ml-6 w-[100%] items-start justify-center">
          <Text className="text-xl font-bold">Status</Text>
        </View>
        <View className="w-[80%] mt-2">
          <View className="ml-6 mb-8">
            {content}
          </View>

          <View className="flex-row items-center ml-6">
            <Flag height={15} width={15} />
            <Text className="font-bold text-xl ml-2">Interessados</Text>
            <Text className="text-green-10">   ver todos &gt;</Text>
          </View>
        </View>
        {interested && interested.interested_users && interested.interested_users.length > 0 ? (
          interested.interested_users.map((user) => (
            <View key={user.id} className="h-20 bg-gray-100 items-center flex-row">
              <CeoAvatar size="w-12 h-12" link={user.ceo_photo} /> {/* Ajuste conforme a estrutura de dados */}
              <View className="ml-2">
                <Text>{user.name}</Text>
                <Text className="text-gray-400">Free Carbon</Text>
              </View>
            </View>
          ))
        ) : (
          <Text className="text-gray-3 font-medium text-[18px] self-center mt-28">Sem interessados</Text>
        )}
      </View>
    </View>
  );
}
