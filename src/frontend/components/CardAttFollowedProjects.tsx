import React from 'react';
import { Text, View, TouchableOpacity, Dimensions } from 'react-native';
import moment from 'moment';
import StarIcon from '../assets/star.svg';
import { styled } from 'nativewind';

const window = Dimensions.get('window');

const CardAttFollowedProject = ({ title, news, data }) => {
  return (
    <View className={`px-4 py-4 bg-white rounded-[20px] my-2`} 
    style={{ 
      marginTop: window.height * 0.01,
      marginBottom: window.height * 0.01,
      width: window.width * 0.9,
      shadowColor: 'rgba(0, 0, 0, 0.60)',
      shadowOffset: { width: 0, height: 4 },
      shadowOpacity: 0.25,
      shadowRadius: 8,
      elevation: 4,
      backgroundColor: '#F6F6F6',
      borderRadius: 20,
    }}>
      <View className="flex flex-row">
        <View className="flex flex-col items-center justify-center ">
          <StarIcon width={25} height={25} />
        </View>
          <View className="ml-4 flex flex-col w-[90%] justify-center">
            <Text className="text-[18px] font-medium mb-2">{title}</Text>
            <Text className="text-[14px] font-semibold text-gray-500 mb-2">{news}</Text>
            <View className="items-end mb-1.5">
            </View>
          </View>
      </View>
    </View>
  );
}

const formatDate = (dateString) => {
  const date = moment(dateString, 'DD/MM/YYYY');
  const today = moment();
  const yesterday = moment().subtract(1, 'days');

  if (date.isSame(today, 'day')) {
    return 'Hoje';
  } else if (date.isSame(yesterday, 'day')) {
    return 'Ontem';
  } else {
    return date.format('DD/MM/YYYY');
  }
};

export default CardAttFollowedProject;
