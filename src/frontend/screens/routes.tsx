import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { createStackNavigator } from '@react-navigation/stack';
import { CommonActions, NavigationContainer, getFocusedRouteNameFromRoute } from '@react-navigation/native';
import { MaterialCommunityIcons } from '@expo/vector-icons';
import React, { useEffect } from 'react';
import { useNavigation, useRoute } from '@react-navigation/native';
import Home from './Home';
import Profile from './Profile';
import Explore from './Explore';
import Notification from './Notification';
import AddProject from './AddProject';
import Opening from './Opening';
import Login from './Login';
import Register from './Register';
import EditProfile from './EditProfile';
import Sinergy from './Sinergy';
import RegisterActivity from './RegisterActivity';
import RegisterEmail from './RegisterEmail';
import RegisterPassword from './RegisterPassword';
import RequestSinergy from './RequestSinergy';
import Project from './Project';
import { Button, View } from 'react-native';
import Map from './Map';
import EditProject from './EditProject';

const Tab = createBottomTabNavigator();
const Stack = createStackNavigator<RootStackParamList>();

export type RootStackParamList = {
  AuthStack: undefined;
  App: undefined;
  ProjectStack: { screen: string, params: { id: number } };
  Request: { projectId: number };
  Início: undefined;
  Project: { id: number };
  RequestSinergy: undefined;
  Login: undefined;
  Cadastro: undefined;
  CadastroEmail: undefined;
  CadastroSenha: undefined;
  Register: undefined;
  Home: undefined;
  Explore: undefined;
  Map: undefined;
  AddProject: undefined;
  Notification: undefined;
  Perfil: undefined;
  Profile: undefined;
  EditProfile: undefined;
  Sinergia: { id: number };
  EditProject: { id: number };
  RegisterActivity: undefined;
  SinergyStack: { screen: string, params: { id: number } };
};

//Stack de telas de autenticação:
function AuthStack() {
  return (
    <Stack.Navigator>
      <Stack.Screen name="Início" component={Opening} options={{ headerShown: false }} />
      <Stack.Screen name="Login" component={Login} />
      <Stack.Screen name="Cadastro" component={Register} />
      <Stack.Screen name="CadastroEmail" component={RegisterEmail} />
      <Stack.Screen name="CadastroSenha" component={RegisterPassword} />
      <Stack.Screen name="Register" component={Register} options={{ title: 'Cadastro' }} />
    </Stack.Navigator>
  );
}

//Stack de telas do perfil:
function ProfileStack() {
  return (
    <Stack.Navigator screenOptions={{ animationEnabled: false }}>
      <Stack.Screen name="Profile" component={Profile} options={{ headerShown: false }} />
      <Stack.Screen name="EditProfile" component={EditProfile} />
      <Stack.Screen name="SinergyStack" component={SinergyStack} options={{ headerShown: false }} />
    </Stack.Navigator>
  );
}

//Stack de telas da sinergia:
function SinergyStack() {
  return (
    <Stack.Navigator screenOptions={{ animationEnabled: false }}>
      <Stack.Screen name="Sinergia" component={Sinergy} />
      <Stack.Screen name="RegisterActivity" component={RegisterActivity} />
    </Stack.Navigator>
  );
}

//Stack de telas do projeto:
function ProjectStack() {
  return (
    <Stack.Navigator>
      <Stack.Screen name="Project" component={Project} options={{ headerShown: false }} />
      <Stack.Screen name="RequestSinergy" component={RequestSinergy} options={{ title: 'Pedido de Sinergia' }} />
      <Stack.Screen name="EditProject" component={EditProject} options={{ title: 'Editar Projeto' }} />
    </Stack.Navigator>
  );
}

//Stack de telas de explorar:
function ExploreStack() {
  return (
    <Stack.Navigator screenOptions={{ animationEnabled: false }}>
      <Stack.Screen name="Explore" component={Explore} options={{ headerShown: false }} />
      <Stack.Screen name="Map" component={Map} options={{ title: 'Mapa' }} />
    </Stack.Navigator>
  );
}


//Abas das telas principais:
function AppTabs() {

  return (
    <Tab.Navigator>
      <Tab.Screen name="Home"
        component={Home}
        options={{
          headerShown: false,
          tabBarLabel: () => null,
          tabBarIcon: ({ focused }) => {
            if (focused) {
              return <MaterialCommunityIcons name='home' size={28} color={'#3A8A88'} />;
            } else {
              return <MaterialCommunityIcons name='home' size={28} color={'#B3B3B3'} />;
            }
          }
        }} />
      <Tab.Screen
        name="Explorar"
        component={ExploreStack}
        options={({ route, navigation }) => {
          const routeName = getFocusedRouteNameFromRoute(route) ?? 'Explore';
          return {
            title: 'Explorar',
            tabBarLabel: () => null,
            tabBarStyle: { display: routeName === 'Map' ? 'none' : 'flex' },
            headerShown: routeName !== 'Map',
            tabBarIcon: ({ focused }) => (
              focused
                ? <MaterialCommunityIcons name='magnify' size={28} color={'#3A8A88'} />
                : <MaterialCommunityIcons name='magnify' size={28} color={'#B3B3B3'} />
            ),
            headerRight: routeName === 'Explore' ? () => (
              <View className="mr-2 rounded">
                <Button
                  onPress={() => navigation.navigate('Map')}
                  title="Mapa"
                  color="#3A8A88"
                />
              </View>
            ) : null,
          };
        }
        }
      />
      <Tab.Screen name="AddProject"
        component={AddProject}
        options={{
          title: 'Adicionar Projeto',
          tabBarLabel: () => null,
          tabBarIcon: ({ focused }) => {
            if (focused) {
              return <MaterialCommunityIcons name='plus-circle-outline' size={28} color={'#3A8A88'} />;
            } else {
              return <MaterialCommunityIcons name='plus-circle-outline' size={28} color={'#B3B3B3'} />;
            }
          }
        }} />
      <Tab.Screen name="Notification"
        component={Notification}
        options={{
          title: 'Notificações',
          tabBarLabel: () => null,
          tabBarIcon: ({ focused }) => {
            if (focused) {
              return <MaterialCommunityIcons name='bell' size={28} color={'#3A8A88'} />;
            } else {
              return <MaterialCommunityIcons name='bell-outline' size={28} color={'#B3B3B3'} />;
            }
          }
        }} />
      <Tab.Screen
        name="Perfil"
        component={ProfileStack}
        options={({ route }) => {
          let routeName = getFocusedRouteNameFromRoute(route);
          if (routeName === undefined) {
            routeName = 'Perfil';
          }
          return {
            headerShown: false,
            tabBarLabel: () => null,
            tabBarIcon: ({ focused }) => {
              if (focused) {
                return <MaterialCommunityIcons name='account' size={28} color={'#3A8A88'} />;
              } else {
                return <MaterialCommunityIcons name='account' size={28} color={'#B3B3B3'} />;
              }
            }
          }
        }}
        listeners={({ navigation }) => ({
          tabPress: event => {
            event.preventDefault();
            navigation.dispatch(
              CommonActions.reset({
                index: 0,
                routes: [{ name: 'Perfil' }],
              }),
            );
          },
        })}
      />
    </Tab.Navigator>
  );
}

function MainStack() {
  return (
    <Stack.Navigator>
      <Stack.Screen name="AuthStack" component={AuthStack} options={{ headerShown: false }} />
      <Stack.Screen name="App" component={AppTabs} options={{ headerShown: false }} />
      <Stack.Screen name="ProjectStack" component={ProjectStack} options={({ route }) => {
        let routeName = getFocusedRouteNameFromRoute(route);
        if (routeName === undefined) {
          routeName = 'ProjectStack';
        }
        return {
          headerShown: false,
        }
      }} />
    </Stack.Navigator>
  );
}

//Exporta as rotas
export default function Routes() {

  return (
    <NavigationContainer >
      <MainStack />
    </NavigationContainer>
  );
}
