import React, { useState } from "react";
import { Text, View, TouchableOpacity, ScrollView, StyleSheet, Dimensions } from "react-native";
import Icon from 'react-native-vector-icons/MaterialIcons';
import Dropdown from "../components/Dropdown";
import CardSinergyProjects from "../components/CardSinergyProjects";
import ActivityCard from "../components/ActivityCard"; // Importa o novo componente
import { useRoute, useNavigation } from '@react-navigation/native';
import { RouteProp } from '@react-navigation/native';
import { RootStackParamList } from './routes';

const window = Dimensions.get('window');

const cards = [
  {
    project: "Way Carbon",
    name: "Guilherme Vasconcellos",
    theme: "Carbonística",
    category: "conservation",
    email: "guilherme@example.com",
    data: "25/05/2024",
    texto: "Pedro, Estou interessado em explorar oportunidades de colaboração para impulsionar a descarbonização. Gostaria de discutir mais sobre como nossas empresas podem trabalhar em conjunto nesse importante objetivo."
  },
  {
    project: "Way Carbon",
    name: "Guilherme Vasconcellos",
    theme: "Carbonística",
    category: "conservation",
    email: "guilherme@example.com",
    data: "25/05/2024",
  },
];

const activities = [
  {
    date: "02 de outubro de 2024",
    title: "Reunião geral",
    items: [
      "Discussão sobre metas de descarbonização.",
      "Planejamento de próximos passos em colaboração com outros CEOs."
    ]
  },
  {
    date: "02 de outubro de 2024",
    title: "Reunião geral",
    items: [
      "Discussão sobre metas de descarbonização.",
      "Planejamento de próximos passos em colaboração com outros CEOs."
    ]
  }
  // ... outras atividades
];

export default function Sinergy({ navigation }) {
  const [selectedThemeValue, setSelectedThemeValue] = useState("");

  const options_themes = [
    { label: 'Negociação', value: '1' }, 
    { label: 'Desenvolvimento', value: '2' },
    { label: 'Finalizada', value: '3' },
  ];
  const route = useRoute<RouteProp<RootStackParamList, 'SinergyStack'>>();  
  const { id } = route.params as unknown as { id: number };
  console.log('id', id);

  return (
    <ScrollView style={{ flex: 1, backgroundColor: '#FFFFFF' }}>
      <View style={{ padding: window.height * 0.025 }}>
        <Text style={{ fontSize: 14, fontWeight: 'bold', marginBottom: 5 }}>Projetos Envolvidos na Integração</Text>
        {cards.map((card, index) => (
          <CardSinergyProjects key={index} {...card} />
        ))}
        <Dropdown 
          type="grayBlack"
          label="Status da Sinergia" 
          options={options_themes} 
          selectedValue={selectedThemeValue} 
          onValueChange={setSelectedThemeValue} 
          placeholder="Negociação" 
        />
        <Text style={{ fontSize: 14, fontWeight: 'bold', marginBottom: 5 }}>Registro de Atividades</Text>
        <TouchableOpacity style={styles.addButton} onPress={() => navigation.navigate('RegisterActivity')}>
          <Icon name="add" size={24} color="#787878" />
          <Text style={styles.addButtonText}>Adicionar novo registro</Text>
        </TouchableOpacity>
        {activities.map((activity, index) => (
          <ActivityCard key={index} {...activity} />
        ))}
      </View>
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  addButton: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    padding: 10,
    marginVertical: 10,
    borderWidth: 3,
    borderColor: '#ccc',
    borderRadius: 5,
    borderStyle: 'dashed',
  },
  addButtonText: {
    marginLeft: 5,
    fontSize: 16,
    color: '#787878',
    fontWeight: 'bold',
  }
});
