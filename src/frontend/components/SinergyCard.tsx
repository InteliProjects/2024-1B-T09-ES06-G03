import { useState } from 'react';
import { View, Text, StyleSheet, Pressable } from 'react-native';
import { useNavigation } from '@react-navigation/native';
import SynergyIcon from '../assets/hands.svg';
import InterestedIcon from '../assets/flag.svg';
import Category, { categoryMap } from './Category';
import CeoAvatar from './CeoAvatar';
import { NavigationProp } from '@react-navigation/native';
import { RootStackParamList } from '../screens/routes';
import GreenNotificationIcon from '../assets/GreenNotificationIcon.svg';

export interface SinergyCardProps {
  id: number;
  targetProjectName: string;
  targetCeoName: string;
  targetProfileAvatar: string;
  targetCompanyName: string;
  targetDescription: string;
  targetCategory: string;
  targetSubcategory: string;
  targetInterestedNumber: number;
  targetSynergyNumber: number;
  sourceProjectName: string;
  status: string;
}

export const sinergyStatusColor = {
  'Em andamento': {titleColor: 'text-[#3A8A88]', bgColor: 'bg-[#C5DAD9]'},
  'Finalizado': {titleColor: 'text-[#BB3756]', bgColor: 'bg-[#E4BBC5]'},
  'Solicitado': {titleColor: 'text-[#4A4444]', bgColor: 'bg-[#BFBDBD]'},
};

export default function SinergyCard({ id, targetProjectName, targetCeoName, targetProfileAvatar, targetCompanyName, targetDescription, targetCategory, targetSubcategory, targetInterestedNumber, targetSynergyNumber, sourceProjectName, status}: SinergyCardProps) {

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
    navigation.navigate('SinergyStack', {
      screen: 'Sinergia',
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
        {/*Header: Notificação*/}
        <View className='flex flex-row mb-3'>
          <View className='flex flex-row items-center w-[50%]' style={{gap: 5}}>
            <GreenNotificationIcon/>
            <Text numberOfLines={1} className='text-[10px] text-green-10 font-medium'>Integrado com {sourceProjectName}</Text>
          </View>
          <View className='flex w-[50%] items-end'>
            <View className={`flex rounded-3xl ${sinergyStatusColor[status].bgColor} items-center justify-center w-[80%] py-[2px]`} >
              <Text className={`text-center font-medium ${sinergyStatusColor[status].titleColor}`}>{status}</Text>
            </View>
          </View>
        </View>

        {/*Corpo: foto, informações e descrição*/}
        <View className="flex flex-row">
          <View className="w-[15%]">
            <CeoAvatar size="w-10 h-10" link={targetProfileAvatar} />
          </View>
          <View className="w-[85%] h-[100%]">
            <View>
              <Text numberOfLines={1} className="text-[18px] font-medium">{targetProjectName}</Text>
              <Text numberOfLines={1} className="text-gray-4 font-medium">{targetCeoName}</Text>
              <Text numberOfLines={1} className="text-gray-4">{targetCompanyName}</Text>
            </View>
            <Text numberOfLines={3} className="my-3 text-[14px]" style={{ height: 14 * 1.4 * 3 }}>{targetDescription}</Text>
          </View>
        </View>

        {/* Footer */}
        <View className="flex flex-row justify-between w-[100%]">

          {/*Categoria e subcategoria*/}
          <View className="flex flex-row items-center w-[60%]">
            <Category category={targetCategory} iconSize={18} circleSize={'w-7 h-7'} />
            <View className={`flex flex-grow ml-3 justify-center rounded-3xl px-3 h-[20px] max-w-[75%] ${categoryMap[targetCategory].color}`}>
              <Text numberOfLines={1} className="text-[12px] text-center text-white font-[500]">{targetSubcategory}</Text>
            </View>
          </View>

          {/* Interessados e Sinergia */}
          <View className="flex flex-row w-[40%] justify-end" style={{ gap: 15 }}>
            <View className="flex flex-row items-center">
              <SynergyIcon width={20} height={20} />
              <Text numberOfLines={1} className="text-gray-3 ml-1">{formatNumber(targetSynergyNumber)}</Text>
            </View>
            <View className="flex flex-row items-center">
              <InterestedIcon width={13} height={13} />
              <Text className="text-gray-3 ml-1">{formatNumber(targetInterestedNumber)}</Text>
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
