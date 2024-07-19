import React from 'react';
import { View, Text, Image, useWindowDimensions } from 'react-native';
import * as Animatable from 'react-native-animatable';

type CarouselItemProps = {
    item: {
        id: number;
        text: string;
        image: any;
    };
};

// Itens do carrossel
export default function CarouselItem({ item }: CarouselItemProps) {
    const { width } = useWindowDimensions();

    return (
        <Animatable.View  animation="fadeIn"  duration={500}  easing="ease-in"  className="flex items-center h-[400px]"  style={{ width }} >
            <Image source={item.image} className="h-[400px]" style={{ width, resizeMode: 'contain' }} />
            <View className="flex h-[500%] w-[100%] items-center p-6">
                <Text className="text-[18px]">{item.text}</Text>
            </View>
        </Animatable.View>
    );
}
