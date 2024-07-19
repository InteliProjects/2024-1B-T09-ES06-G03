import React from 'react';
import { Text, StyleSheet, TouchableHighlight, GestureResponderEvent } from 'react-native';
import { LinearGradient } from 'expo-linear-gradient';
import { styled } from 'nativewind';

const StyledLinearGradient = styled(LinearGradient);

interface ShadowButtonProps {
  onPress: any;
  title: string;
}

const ShadowButton: React.FC<ShadowButtonProps> = ({ onPress, title }) => {
  return (
    <TouchableHighlight onPress={onPress} className='rounded-2xl p-3 bg-white' style={[styles.buttonShadow]} underlayColor="#d9d9d9">
        <Text className='text-black text-[18px] font-medium text-center'>
          {title}
        </Text>
    </TouchableHighlight>
  );
};

const styles = StyleSheet.create({
  buttonShadow: {
    shadowColor: '#000',
    shadowOffset: {
      width: 6,
      height: 6,
    },
    shadowOpacity: 0.80,
    shadowRadius: 4,
    elevation: 4,
  },
});

export default ShadowButton;
