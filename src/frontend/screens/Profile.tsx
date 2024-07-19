import React, { useEffect, useState } from "react";
import { Text, View, StyleSheet, ImageBackground, TouchableOpacity, FlatList, Linking } from "react-native";
import CeoAvatar from "../components/CeoAvatar";
import Svg, { Circle, G, LinearGradient, Path, Stop } from "react-native-svg";
import ProjectCard from "../components/ProjectCard";
import SinergyCard from "../components/SinergyCard";
import { ceoApi, projectApi, coreApi } from "../services/api";
import Loading from "../components/Loading";

export interface profileData {
  id?: number,
  name?: string,
  email?: string,
  company?: string,
  instagram?: string,
  linkedin?: string,
  photo?: string,
  description?: string
}

const sinergiaData = [

  {
    id: 1,
    targetProjectName: 'Projeto 2',
    targetCeoName: 'Davi',
    targetProfileAvatar: 'https://media.licdn.com/dms/image/D5603AQEkp6Rjz7aZLQ/profile-displayphoto-shrink_200_200/0/1675300303921?e=1721865600&v=beta&t=Er53SeGHN3Yw3QptOWO2TgJpXr7OximYg99mt3OxEB0',
    targetCompanyName: `Davi's`,
    targetDescription: 'Projeto com o intuito de um mundo melhor',
    targetCategory: 'productivity',
    targetSubcategory: 'Gestão',
    targetInterestedNumber: 100,
    targetSynergyNumber: 100,
    sourceProjectName: 'Projeto 1',
    status: 'Em andamento'
  }
]

export default function Profile({ navigation }) {
  const [boolProjeto, setBoolProjeto] = useState(true)
  const [userId, setUserId] = useState(0)
  const [projects, setProjects] = useState([])
  const [profileData, setProfileData] = useState<profileData>({})
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchUserProjects = async () => {
      try{
        const response = await projectApi(`/projects/me`);
        setUserId(response.data[0].user_id)
      } catch (error) {
        console.error('Error fetching projetos:', error);
      }
    }

    const fetchProfileData = async () => {
      try {
        const response = await coreApi(`/user/${userId}`)
        setProfileData(response.data)
      } catch (error) {
        console.log('Ocorreu um erro ao buscar as informações do usuário')
      }
    }

    const fetchProjectData = async () => {
      try {
        const response = await projectApi(`/projects`);
        const userProjects = response.data.filter(item => item.user_id === userId)
        setProjects(userProjects)
        setLoading(false)

      } catch (error) {
        alert('Ocorreu um erro ao buscar as informações dos projetos')
      }
    };

    fetchUserProjects();
    fetchProfileData();
    fetchProjectData();

  }, [userId]);

  const linkExterno = async (url: string) => {
    // Verifica se o link pode ser aberto
    const supported = await Linking.canOpenURL(url);

    if (supported) {
      // Se o link é suportado, abre o link
      await Linking.openURL(url);
    } else {
      console.log(`Don't know how to open this URL: ${url}`);
    }
  }




  if (loading) {
    return (
      <View className="w-[100%] h-[100%] justify-center items-center">
        <Loading />
      </View>
    )
  } else {

    return (
      <View style={styles.divPrincipal}>
        <View className='h-[25%]'>
          <ImageBackground
            source={{ uri: 'https://media.licdn.com/dms/image/D4E16AQF6Bi9B-lNJLA/profile-displaybackgroundimage-shrink_350_1400/0/1678717278492?e=1721865600&v=beta&t=u6Gmh3y0d2_sKFG3NgKUPLazxo1rFS5LGlDJD_pkIWE' }}
            style={styles.imagemBackgroud}
            resizeMode="cover"
          >
            <TouchableOpacity style={styles.touchEditar} onPress={() => navigation.navigate('EditProfile')}>
              <Text style={styles.textEditar}>Editar</Text>
            </TouchableOpacity>
          </ImageBackground>
        </View>
        <View style={styles.divInfos}>
          <View style={styles.divImageSocial}>
            <CeoAvatar size="w-32 h-32" link={profileData.photo} />
            <View style={styles.divIconSocial}>
              {profileData.linkedin &&
                <TouchableOpacity onPress={() => linkExterno(profileData.linkedin)}>
                  <Svg width="30" height="30" viewBox="0 0 25 25" fill="none">
                    <Circle cx="12.4031" cy="12.4031" r="12.4031" fill="#0076B2" />
                    <Path d="M5.6862 10.0737H8.56529V19.349H5.6862V10.0737ZM7.12654 5.45752C7.45676 5.45752 7.77956 5.55558 8.0541 5.73931C8.32864 5.92303 8.5426 6.18416 8.66889 6.48965C8.79519 6.79515 8.82815 7.13129 8.76362 7.45554C8.69908 7.7798 8.53994 8.07761 8.30633 8.31129C8.07272 8.54497 7.77513 8.70403 7.45122 8.76834C7.12731 8.83265 6.79163 8.79932 6.48664 8.67258C6.18165 8.54583 5.92105 8.33136 5.73783 8.05629C5.5546 7.78123 5.45696 7.45793 5.45728 7.1273C5.4577 6.6843 5.63375 6.25959 5.94675 5.94649C6.25976 5.63339 6.6841 5.45752 7.12654 5.45752ZM10.3713 10.0737H13.1311V11.3471H13.1693C13.554 10.6181 14.492 9.84924 15.8926 9.84924C18.8082 9.84288 19.3487 11.7642 19.3487 14.2553V19.349H16.4697V14.8363C16.4697 13.7618 16.4506 12.3786 14.9737 12.3786C13.4968 12.3786 13.2456 13.5501 13.2456 14.7663V19.349H10.3713V10.0737Z" fill="white" />
                  </Svg>
                </TouchableOpacity>
              }
              {profileData.instagram &&
                <TouchableOpacity onPress={() => linkExterno(profileData.instagram)}>
                  <Svg height="30" width="30" viewBox="0 0 128 128">
                    <LinearGradient
                      id="a"
                      gradientTransform="matrix(1 0 0 -1 594 633)"
                      gradientUnits="userSpaceOnUse"
                      x1="-566.711"
                      x2="-493.288"
                      y1="516.569"
                      y2="621.43"
                    >
                      <Stop offset="0" stopColor="#ffb900" />
                      <Stop offset="1" stopColor="#9100eb" />
                    </LinearGradient>
                    <Circle cx="64" cy="64" fill="url(#a)" r="64" />
                    <G fill="#fff">
                      <Path d="M82.333 104H45.667C33.72 104 24 94.281 24 82.333V45.667C24 33.719 33.72 24 45.667 24h36.666C93.281 24 103 33.719 103 45.667v36.666C103 94.281 93.281 104 82.333 104zM45.667 30.667c-8.271 0-15 6.729-15 15v36.667c0 8.271 6.729 15 15 15h36.666c8.271 0 15-6.729 15-15V45.667c0-8.271-6.729-15-15-15z" />
                      <Path d="M64 84c-11.028 0-20-8.973-20-20 0-11.029 8.972-20 20-20s20 8.971 20 20c0 11.027-8.972 20-20 20zm0-33.333c-7.352 0-13.333 5.981-13.333 13.333 0 7.353 5.981 13.333 13.333 13.333s13.333-5.98 13.333-13.333c0-7.352-5.98-13.333-13.333-13.333z" />
                      <Circle cx="85.25" cy="42.75" r="4.583" />
                    </G>
                  </Svg>
                </TouchableOpacity>
              }
            </View>
          </View>
          <View style={styles.divDescricaoInicial}>
            <View style={styles.divNome}>
              <Text style={styles.textNome}>{profileData.name}</Text>
              <Text style={styles.textEmpresa}>{profileData.company}</Text>
            </View>
            <View style={styles.divProjetosInteressados}>
              <View style={styles.divFilhoProjetoInterresados}>
                <Text style={styles.textNumeroBold}>{projects.length}</Text>
                <Text>Projetos</Text>
              </View>
              <View style={styles.divFilhoProjetoInterresados}>
                <Text style={styles.textNumeroBold}>35</Text>
                <Text>Interresados</Text>
              </View>
            </View>
          </View>
          <View style={styles.descricaoGrande}>
            <Text>{profileData.description}</Text>
          </View>
        </View>
        <View style={styles.divProjetosSinergiasButton}>
          <TouchableOpacity onPress={() => setBoolProjeto(true)}>
            <Text style={boolProjeto ? styles.textProjetoSinergiaSelecionado : styles.textProjetoSinergia}>Projetos</Text>
          </TouchableOpacity>
          <TouchableOpacity onPress={() => setBoolProjeto(false)}>
            <Text style={boolProjeto ? styles.textProjetoSinergia : styles.textProjetoSinergiaSelecionado}>Sinergias</Text>
          </TouchableOpacity>
        </View>
        <View style={styles.hr} ></View>
        <View style={styles.indicator} className={`${boolProjeto ? 'self-start ml-[17%]' : 'self-end mr-[18%]'}`}></View>

        {boolProjeto ? (
          <FlatList
            style={styles.divCards}
            data={projects}
            renderItem={({ item }) => (
              <View className="mt-3 mb-3">
                <ProjectCard id={item.id} projectName={item.name} ceoName={item.ceo_name} profileAvatar={item.photo} companyName={item.company_name} description={item.description} category={item.category_name} subcategory={item.subcategory_name} interestedNumber={item.interested_count} synergyNumber={item.synergy_count} />
              </View>
            )}
            keyExtractor={(item) => item.id.toString()}
            contentContainerStyle={styles.divCards}
          />
        ) : (
          <FlatList
            style={styles.divCards}
            data={sinergiaData}
            renderItem={({ item }) => (
              <View className="mt-3 mb-3">
                <SinergyCard
                  key={item.id}
                  id={item.id}
                  targetProjectName={item.targetProjectName}
                  targetCeoName={item.targetCeoName}
                  targetProfileAvatar={item.targetProfileAvatar}
                  targetCompanyName={item.targetCompanyName}
                  targetDescription={item.targetDescription}
                  targetCategory={item.targetCategory}
                  targetSubcategory={item.targetSubcategory}
                  targetInterestedNumber={item.targetInterestedNumber}
                  targetSynergyNumber={item.targetSynergyNumber}
                  sourceProjectName={item.sourceProjectName}
                  status={item.status}
                />
              </View>
            )}
            keyExtractor={(item) => item.id.toString()}
            contentContainerStyle={styles.divCards}
          />
        )}


      </View >
    );
  }
}

const styles = StyleSheet.create({

  divPrincipal: {
    display: 'flex',
    flexDirection: 'column',
    width: '100%',
    height: '100%'
  },

  imagemBackgroud: {
    width: '100%',
    height: '100%',
    justifyContent: 'center',
    alignItems: 'center',
    marginTop: -24
  },

  touchEditar: {
    backgroundColor: '#818181a5',
    borderRadius: 1000,
    alignItems: 'center',
    justifyContent: 'center',
    paddingVertical: 4,
    paddingHorizontal: 25,
    marginLeft: 'auto',
    marginRight: 10
  },

  textEditar: {
    color: 'white',
  },

  divInfos: {
    display: 'flex',
    flexDirection: 'column'
  },

  divImageSocial: {
    display: 'flex',
    flexDirection: 'row',
    width: '100%',
    marginTop: -80,
    marginLeft: 35
  },

  divIconSocial: {
    height: '70%',
    width: '50%',
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'flex-end',
    alignItems: 'flex-end',
    gap: 10,
    marginTop: 10,
    marginLeft: 10
  },

  divDescricaoInicial: {
    display: 'flex',
    flexDirection: 'row',
    width: '100%',
    justifyContent: 'space-around',
    alignItems: 'center',
  },

  divNome: {
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center',
    width: '50%'
  },

  textNome: {
    fontSize: 15,
    fontWeight: 'bold'
  },

  textEmpresa: {
    fontSize: 15,
  },

  divProjetosInteressados: {
    display: 'flex',
    flexDirection: 'row',
    gap: 10,
    padding: 10,
    width: '50%',
    justifyContent: 'space-around'
  },

  divFilhoProjetoInterresados: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
    gap: 5
  },

  textNumeroBold: {
    fontWeight: 'bold'
  },

  descricaoGrande: {
    display: 'flex',
    width: '100%',
    padding: 20
  },

  divProjetosSinergiasButton: {
    display: 'flex',
    flexDirection: 'row',
    width: '100%',
    justifyContent: 'space-evenly',
    alignItems: 'center',
  },

  textProjetoSinergia: {
    fontSize: 20,
  },

  textProjetoSinergiaSelecionado: {
    fontSize: 20,
    fontWeight: "500"
  },

  hr: {
    borderBottomColor: '#B3B3B3',
    borderBottomWidth: 2,
    marginTop: 10,
  },

  indicator: {
    borderBottomColor: '#3A8A88',
    borderBottomWidth: 2,
    width: '25%',
    marginBottom: 20,
    position: 'relative',
    zIndex: 1,
    marginTop: -2,
  },

  divCards: {
    display: 'flex',
    flexDirection: 'column',
  },

})
