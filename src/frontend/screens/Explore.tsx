import React, { useEffect, useState } from "react";
import { Text, View, TouchableOpacity, ScrollView, FlatList, TextInput } from "react-native";
import Category, { categoryMap } from "../components/Category";
import ProjectCard, { subCategoryTemplate } from "../components/ProjectCard";
import { MaterialCommunityIcons } from '@expo/vector-icons';
import { Badge } from "../components/Badge";
import Loading from "../components/Loading";
import { projectApi } from "../services/api";

export default function Explore({ navigation }) {
  const [projects, setProjects] = useState([]);
  const [search, setSearch] = useState('');
  const [categoriaAtual, setCategoriaAtual] = useState(null);
  const [subCategory, setSubCategory] = useState([]);
  const [subCategorySelect, setSubCategorySelect] = useState(null);
  const [dataProjects, setDataProjects] = useState([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await projectApi(`/projects`);
        console.log(response.data)
        setDataProjects(response.data);
        setProjects(response.data)
        setLoading(false)

      } catch (error) {

      }
    };

    fetchData();
  }, []);

  useEffect(() => {
    filterProjects();
  }, [search, categoriaAtual, subCategorySelect]);

  const filterProjects = () => {
    let filteredProjects = dataProjects;

    if (categoriaAtual) {
      filteredProjects = filteredProjects.filter(project => project.category_name === categoriaAtual);
    }
    console.log(filteredProjects)

    if (subCategorySelect) {
      filteredProjects = filteredProjects.filter(project => project.subcategory_name === subCategorySelect);
    }

    if (search) {
      const searchText = search.toUpperCase();
      filteredProjects = filteredProjects.filter(item => {
        const itemProjectName = item.name ? item.name.toUpperCase() : '';
        const itemCeoName = item.user_id ? item.ceo_name.toUpperCase() : '';
        return itemProjectName.includes(searchText) || itemCeoName.includes(searchText);
      });
    }
    setProjects(filteredProjects);
  };

  const handleCategoryFilter = (category) => {
    if (category === categoriaAtual) {
      setCategoriaAtual(null);
      setSubCategory([]);
    } else {
      setCategoriaAtual(category);
      setSubCategory(subCategoryTemplate[category] || []);
    }
    setSubCategorySelect(null); // Reset subcategory selection on category change
  };

  const handleSubCategoryFilter = (subCategoryFunc) => {
    setSubCategorySelect(subCategoryFunc === subCategorySelect ? null : subCategoryFunc);
  };

  const handleSearch = (text) => {
    setSearch(text);
  };

  if (loading) {
    return (
      <View className="w-[100%] h-[100%] justify-center items-center">
        <Loading />
      </View>
    )
  } else {

    return (

      <View className="flex flex-col h-[100%] bg-white">

        {/* Barra de pesquisa */}
        <View className="flex justify-end mx-5 mb-[5%] h-[10%]">
          <View className=' flex flex-row w-[100%] bg-[#ABABAB] py-3 px-3 rounded-full'>
            <View className='w-[15%] ml-[2%]'>
              <MaterialCommunityIcons name='magnify' size={28} color={'black'} />
            </View>
            <TextInput
              className='w-[83%]'
              placeholder="Pesquise um projeto ou CEO..."
              placeholderTextColor="#4c4c4c"
              value={search}
              onChangeText={(text) => handleSearch(text)}
            />
          </View>
        </View>

        <Text className="text-[20px] font-medium mx-5 mb-[5%] h-[5%] ">Temas</Text>

        {/* Scroll de categorias */}
        <View className="h-[12%] mb-[5%]">
          <ScrollView horizontal={true} showsHorizontalScrollIndicator={false}>
            <TouchableOpacity className="items-center w-[30vw] " onPress={() => handleCategoryFilter('conservation')}>
              <Category category={'conservation'} circleSize={'w-[50px] h-[50px]'} iconSize={20} />
              <Text className={`text-[12px] text-center ${categoriaAtual == 'conservation' && 'font-medium'}`}>Conservação do Planeta</Text>
            </TouchableOpacity>
            <TouchableOpacity className="items-center w-[30vw]" onPress={() => handleCategoryFilter('productivity')}>
              <Category category={'productivity'} circleSize={'w-[50px] h-[50px]'} iconSize={20} />
              <Text className={`text-[12px] text-center ${categoriaAtual == 'productivity' && 'font-medium'}`}>Produtividade e Competitividade</Text>
            </TouchableOpacity>
            <TouchableOpacity className="items-center w-[30vw]" onPress={() => handleCategoryFilter('health')}>
              <Category category={'health'} circleSize={'w-[50px] h-[50px]'} iconSize={20} />
              <Text className={`text-[12px] text-center ${categoriaAtual == 'health' && 'font-medium'}`}>Bem Estar, Saúde e Felicidade</Text>
            </TouchableOpacity>
            <TouchableOpacity className="items-center w-[30vw]" onPress={() => handleCategoryFilter('diversity')}>
              <Category category={'diversity'} circleSize={'w-[50px] h-[50px]'} iconSize={20} />
              <Text className={`text-[12px] text-center ${categoriaAtual == 'diversity' && 'font-medium'}`}>D&I Dignidade e Integridade</Text>
            </TouchableOpacity>
            <TouchableOpacity className="items-center w-[30vw]" onPress={() => handleCategoryFilter('environmentalImpact')}>
              <Category category={'environmentalImpact'} circleSize={'w-[50px] h-[50px]'} iconSize={20} />
              <Text className={`text-[12px] text-center ${categoriaAtual == 'environmentalImpact' && 'font-medium'}`}>Impacto Ambiental</Text>
            </TouchableOpacity>
            <TouchableOpacity className="items-center w-[30vw]" onPress={() => handleCategoryFilter('integrity')}>
              <Category category={'integrity'} circleSize={'w-[50px] h-[50px]'} iconSize={20} />
              <Text className={`text-[12px] text-center ${categoriaAtual == 'integrity' && 'font-medium'}`}>Integridade</Text>
            </TouchableOpacity>
          </ScrollView>
        </View>

        {subCategory.length > 0 && (
          <>
            <Text className="text-[20px] font-medium mx-5 mb-[5%]">Subtemas</Text>
            <View className="flex justify-start items-start w-[100%] mb-[5%] mx-3">
              <FlatList
                className="w-[100%]"
                horizontal={true}
                showsHorizontalScrollIndicator={false}
                data={subCategory}
                keyExtractor={(item) => item}
                renderItem={({ item }) => (
                  <TouchableOpacity className="items-center pr-2" onPress={() => handleSubCategoryFilter(item)}>
                    <Badge label={item} labelClasses="text-white" className={`${categoryMap[categoriaAtual].color}`} variant="default" />
                  </TouchableOpacity>
                )}
              />
            </View>
          </>
        )}

        {/* Lista de projetos */}
        <View className="flex justify-center items-center w-[100%] h-[68%]">
          <FlatList
            className={`w-[100%] h-[100%] ${subCategory.length > 0 ? 'mb-[28%]' : 'mb-[5%]'}`}
            data={projects}
            keyExtractor={(item) => item.id.toString()}
            renderItem={({ item }) => (
              <View className="mt-3 mb-3">
                <ProjectCard id={item.id} projectName={item.name} ceoName={item.ceo_name} profileAvatar={item.photo} companyName={item.company_name} description={item.description} category={item.category_name} subcategory={item.subcategory_name} interestedNumber={item.interested_count} synergyNumber={item.synergy_count} />
              </View>
            )}
          />
        </View>
      </View>
    );
  }
}