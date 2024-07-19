import { useState } from 'react';
import { View, Text, StyleSheet, Pressable } from 'react-native';
import { useNavigation } from '@react-navigation/native';
import SynergyIcon from '../assets/hands.svg';
import InterestedIcon from '../assets/flag.svg';
import Category, { categoryMap } from './Category';
import CeoAvatar from './CeoAvatar';
import { NavigationProp } from '@react-navigation/native';
import { RootStackParamList } from '../screens/routes';

export const subCategoryTemplate: any = {
  conservation: ["Cultivo e Regeneração"],
  productivity: ["Capacitar para Ampliar Acesso"],
  health: ["Bem-Estar e Saúde Mental"],
  diversity: ["Raça", "Mulheres", "DE&I Geral"],
  environmentalImpact: ["Descarbonização", "Economia Circular"],
  integrity: []
}

export interface ProjectCardProps {
  id: number;
  projectName: string;
  ceoName: string;
  profileAvatar: string;
  companyName: string;
  description: string;
  category: string;
  subcategory: string;
  interestedNumber: number;
  synergyNumber: number;
}

export default function ProjectCard({
  id,
  projectName,
  ceoName,
  profileAvatar,
  companyName,
  description,
  category,
  subcategory,
  interestedNumber,
  synergyNumber
}: ProjectCardProps) {

  const formatNumber = (num) => {
    if (num < 1000) {
      return num.toString();
    } else if (num == undefined || num == null || num == 0) {
      return '0';
    } else {
      return (num / 1000).toFixed(1) + 'k';
    }
  };

  const [bgColor, setBgColor] = useState('bg-gray-1');
  const navigation = useNavigation<NavigationProp<RootStackParamList>>();

  const handlePress = () => {
    navigation.navigate('ProjectStack', {
      screen: 'Project',
      params: { id: id }
    });
  };

  return (
    <Pressable
      onPressIn={() => setBgColor('bg-gray-2')}
      onPressOut={() => setBgColor('bg-gray-1')}
      style={[styles.shadowCard]}
      onPress={handlePress}
    >
      <View className={`flex flex-col justify-center ${bgColor} w-[330px] rounded-2xl px-4 py-4 self-center`} style={[styles.shadowCard]}>
        
        {/*Corpo: foto, informações e descrição*/}
        <View className="flex flex-row">
          <View className="w-[15%]">
            <CeoAvatar size="w-10 h-10" link={profileAvatar} />
          </View>
          <View className="w-[85%] h-[100%]">
            <View>
              <Text numberOfLines={1} className="text-[18px] font-medium">{projectName}</Text>
              <Text numberOfLines={1} className="text-gray-4 font-medium">{ceoName}</Text>
              <Text numberOfLines={1} className="text-gray-4">{companyName}</Text>
            </View>
            <Text numberOfLines={3} className="my-3 text-[14px]" style={{ height: 14 * 1.4 * 3 }}>{description}</Text>
          </View>
        </View>

        {/* footer */}
        <View className="flex flex-row justify-between w-[100%]">

          {/* Categoria e subcategoria */}
          <View className="flex flex-row items-center w-[60%]">
            <Category category={category} iconSize={18} circleSize={'w-7 h-7'} />
            <View className={`flex flex-grow ml-3 justify-center rounded-3xl px-3 h-[20px] max-w-[75%] ${categoryMap[category].color}`}>
              <Text numberOfLines={1} className="text-[12px] text-center text-white font-[500]">{subcategory}</Text>
            </View>
          </View>

          {/* Números de sinergias e interessados */}
          <View className="flex flex-row w-[40%] justify-end" style={{ gap: 15 }}>
            <View className="flex flex-row items-center">
              <SynergyIcon width={20} height={20} />
              <Text numberOfLines={1} className="text-gray-3 ml-1">{formatNumber(synergyNumber)}</Text>
            </View>
            <View className="flex flex-row items-center">
              <InterestedIcon width={13} height={13} />
              <Text className="text-gray-3 ml-1">{formatNumber(interestedNumber)}</Text>
            </View>
          </View>
        </View>
      </View>
    </Pressable>
  );
}

const styles = StyleSheet.create({
  shadowCard: {
    shadowColor: 'rgba(0, 0, 0, 0.60)',
    shadowOffset: { width: 0, height: 4 },
    shadowOpacity: 0.25,
    shadowRadius: 8,
    elevation: 4,
  },
});
