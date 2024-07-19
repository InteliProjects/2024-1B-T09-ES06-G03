import React, { useState, useRef } from 'react';
import { View, FlatList, Animated, ViewToken, ViewabilityConfig } from 'react-native';
import CarouselItem from './CarouselItem';
import Paginator from './Paginator';

type ContentItem = {
    id: number;
    text: string;
    image: any;
};

// Conteúdo do carrossel
const content: ContentItem[] = [
    {
        id: 1,
        text: 'Conecte-se com líderes visionários',
        image: require('../assets/carousel-1.png')
    },
    {
        id: 2,
        text: 'Fortaleça sua rede profissional',
        image: require('../assets/carousel-2.png')
    },
    {
        id: 3,
        text: 'Potencialize projetos e inove',
        image: require('../assets/carousel-3.png')
    }
];

export default function Carousel() {
    const [currentIndex, setCurrentIndex] = useState<number>(0);
    const scrollX = useRef(new Animated.Value(0)).current;
    const viewableItemsChanged = useRef(({ viewableItems }: { viewableItems: ViewToken[] }) => {
        if (viewableItems.length > 0) {
            setCurrentIndex(viewableItems[0].index || 0);
        }
    }).current;

    const viewConfig: ViewabilityConfig = { viewAreaCoveragePercentThreshold: 50 };
    const slidesRef = useRef<FlatList<ContentItem>>(null);

    return (
        // Flatlist com as propriedades do carrossel
        <View style={{ height: 500, alignItems: 'center', justifyContent: 'center' }}>
            <FlatList
                data={content}
                renderItem={({ item }) => <CarouselItem item={item} />}
                horizontal
                showsHorizontalScrollIndicator={false}
                pagingEnabled
                bounces={false}
                decelerationRate="fast"
                keyExtractor={(item) => item.id.toString()}
                onScroll={Animated.event(
                    [{ nativeEvent: { contentOffset: { x: scrollX } } }],
                    { useNativeDriver: false }
                )}
                scrollEventThrottle={16}
                onViewableItemsChanged={viewableItemsChanged}
                viewabilityConfig={viewConfig}
                ref={slidesRef}
            />
            <Paginator data={content} scrollX={scrollX} />
        </View>
    );
}
