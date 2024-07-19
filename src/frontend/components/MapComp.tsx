import React, { useEffect, useState } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import MapView, { Marker, Callout, PROVIDER_GOOGLE } from 'react-native-maps';
import CeoAvatar from './CeoAvatar';
import { coreApi, projectApi } from '../services/api';

const initialLocation = {
    latitude: -23.5729072,
    longitude: -46.7064059,
    latitudeDelta: 0.0922,
    longitudeDelta: 0.0421,
};

//Endpoint APi /geocode?address={endereço_projeto}
//Response "lat": -23.5729072,"long": -46.70640590000001

export default function MapComp() {
    const [projects, setProjects] = useState([])

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await projectApi('/projects');
                const projectsWithGeo = [];

                for (const project of response.data) {
                    try {
                        const responseGeo = await coreApi(`/geocode?address=${project.local}`);
                        projectsWithGeo.push({
                            ...project,
                            geoLocationLat: responseGeo.data.lat,
                            geoLocationLong: responseGeo.data.long
                        });
                    } catch (error) {
                    }
                    // Atualiza o estado após cada iteração
                    setProjects([...projectsWithGeo]);
                }

                // Atualiza o estado uma última vez
                setProjects(projectsWithGeo);
            } catch (error) {
                console.error('Erro ao atualizar projetos com geolocalização:', error);
            }
        };

        fetchData();
    }, []);

    return (
        <View className='flex w-full h-full'>
            <MapView
                provider={PROVIDER_GOOGLE}
                className='flex w-full h-full'
                initialRegion={initialLocation}
            >
                {projects.map((project) => (
                    <Marker key={project.id} coordinate={{ latitude: project.geoLocationLat, longitude: project.geoLocationLong, }}>
                        <View className='flex items-center'>
                            <CeoAvatar size={'w-10 h-10'} link={project.photo} />
                        </View>
                        <Callout>
                            <View style={styles.calloutContainer}>
                                <Text className='mt-1 font-semibold'>{project.name}</Text>
                                <Text className='font-normal'>Criado por: {project.ceo_name}</Text>
                            </View>
                        </Callout>
                    </Marker>
                ))}
            </MapView>
        </View>
    );
};

const styles = StyleSheet.create({
    calloutContainer: {
        width: 200,
        height: 'auto',
    },
});
