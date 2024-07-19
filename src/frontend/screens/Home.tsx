import { useEffect, useState } from 'react';
import { View, Text, TouchableOpacity, ImageBackground, FlatList, Image } from 'react-native';
import GradientButton from '../components/GradientButton';
import ProjectCard, { ProjectCardProps } from '../components/ProjectCard';
import LogoWhite from '../assets/LogoWhite.svg';
import { authApi, ceoApi, coreApi, projectApi } from '../services/api';
import { profileData } from './Profile';
import Loading from '../components/Loading';
import { set } from 'react-hook-form';

export default function Home({ navigation }) {

  const [recomendProjects, setRecomendProjects] = useState<ProjectCardProps[]>([]);
  const [projects, setProjects] = useState<ProjectCardProps[]>([]);
  const [userIdHome, setUserIdHome] = useState(null)
  const [profileData, setProfileData] = useState<profileData>({})
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchUserProjects = async () => {
      try{
        const response_userID = await projectApi(`/projects/me`);
        setUserIdHome(response_userID.data[0].user_id)

        const response_userData = await authApi(`/user/${response_userID.data[0].user_id}`)
        setProfileData(response_userData.data)

        const response_predict = await projectApi.post(`/projects/predict`, { user_id: response_userID.data[0].user_id });
        setRecomendProjects(response_predict.data);

        const response_projects = await projectApi(`/projects`);
        let projects = response_projects.data.filter(item => item.user_id !== response_userID.data[0].user_id)
        // Ordenando os projetos pela data de criação em ordem decrescente
        projects = projects.sort((a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime());
        setProjects(projects);
        setLoading(false)
      } catch (error) {
        console.error('Error fetching projetos:', error);
        console.log(2)
      }
    }

    fetchUserProjects()
  }, []);

  const renderProjectCard = ({ item }: { item }) => (
    <View className='items-center justify-center ml-5 mr-2'>
      <ProjectCard id={item.id} projectName={item.name} ceoName={item.ceo_name} profileAvatar={item.photo} companyName={item.company_name} description={item.description} category={item.category_name} subcategory={item.subcategory_name} interestedNumber={item.interested_count} synergyNumber={item.synergy_count} />
    </View>
  );

  if (loading) {
    return (
      <View className="w-[100%] h-[100%] justify-center items-center">
        <Loading />
      </View>
    )
  } else {


    return (
      // Página inteira
      <View className="w-[100%] h-[100%] bg-white">
        {/* Header */}
        <View className='h-[30%]'>
          <Image source={require('../assets/backgroundImage.png')} className='h-[100%]' />
          {/* Layer2 */}
          <View className='absolute top-[18%] left-[5%] right-[5%] z-10 items-center gap-4'>
            <LogoWhite />
            <Text className='text-white text-[32px] text-center'>Olá, {profileData.name}!</Text>
            <Text className='text-center text-white text-[16px] w-[80%] self-center'>
              Conecte-se com projetos e CEOs de diversas indústrias e fortaleça sua rede profissional.
            </Text>
            <View className='w-[50%]'>
              <GradientButton onPress={() => navigation.navigate('Explore')} title={'Explorar Projetos'} />
            </View>
          </View>
        </View>

        {/* Corpo */}
        <View className='h-[30%] mt-[10%]'>
          {/* Seção Recomendados */}
          <View className='flex flex-row items-center mb-2'>
            <Text className='text-[18px] font-medium mx-5'>Recomendados para você</Text>
            <Text className='text-[12px] text-green-10 mt-1'>Ver todos {'>'} </Text>
          </View>
          <FlatList
            data={recomendProjects}
            horizontal={true}
            keyExtractor={item => item.id.toString()}
            renderItem={renderProjectCard}
            showsHorizontalScrollIndicator={false}
          />
        </View>

        {/* Seção Novidades */}
        <View className='h-[30%] mt-[5%]'>
          <View className='flex flex-row items-center mb-2'>
            <Text className='text-[18px] font-medium mx-5'>Novidades</Text>
            <Text className='text-[12px] text-green-10 mt-1'>Ver todos {'>'} </Text>
          </View>
          <FlatList
            data={projects}
            horizontal={true}
            keyExtractor={item => item.id.toString()}
            renderItem={renderProjectCard}
            showsHorizontalScrollIndicator={false}
          />
        </View>
      </View>
    );
  }
}
