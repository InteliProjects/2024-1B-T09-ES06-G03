import React from 'react';
import { Text, StyleSheet, TouchableHighlight, GestureResponderEvent } from 'react-native';
import { LinearGradient } from 'expo-linear-gradient';
import { styled } from 'nativewind';

const StyledLinearGradient = styled(LinearGradient);

interface GradientButtonProps {
  onPress: (event: GestureResponderEvent) => void;
  title: string;
}

const GradientButton: React.FC<GradientButtonProps> = ({ onPress, title }) => {
  return (
    <TouchableHighlight onPress={onPress} className='rounded-2xl w-full'>
      <StyledLinearGradient  colors={['#14514F', '#3A8A88']}  start={{ x: 0, y: 0 }}  end={{ x: 1, y: 0 }}  className='p-3 rounded-2xl'  style={[styles.buttonShadow]}>
        <Text className='text-white text-[18px] font-medium text-center'>
          {title}
        </Text>
      </StyledLinearGradient>
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
    elevation: 5,
  },
});

export default GradientButton;
