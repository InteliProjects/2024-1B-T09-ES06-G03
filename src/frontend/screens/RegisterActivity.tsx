import React, { useState } from "react";
import { Text, View, TextInput, TouchableOpacity, Dimensions, StyleSheet } from "react-native";
import CardSinergyProjects from "../components/CardSinergyProjects";
import GradientButton from "../components/GradientButton";
import InputComponent from "../components/Input";
import DateInput from "../components/DateInput"; // Componente personalizado de entrada de data
import Tabs from "../components/Tabs";

const window = Dimensions.get('window');

const cards = [
  {
    project: "Bovaer",
    name: "Pedro Oliveira",
    theme: "Free Carbon",
    category: "conservation",
    email: "pedro@example.com",
    data: "25/05/2024",
    texto: "Pedro, Estou interessado em explorar oportunidades de colaboração para impulsionar a descarbonização. Gostaria de discutir mais sobre como nossas empresas podem trabalhar em conjunto nesse importante objetivo."
  },
  {
    project: "Abraer",
    name: "Igor Alexandre",
    theme: "Going Up",
    category: "conservation",
    email: "igor@example.com",
    data: "25/05/2024",
  },
];

export default function RegisterActivity({ navigation }) {
  const [step, setStep] = useState(0);
  const [date, setDate] = useState("");
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");

  const handleNext = () => {
    if (step < 2) {
      setStep(step + 1);
    } else {
      // Aqui você pode lidar com a submissão do formulário
      console.log("Date:", date);
      console.log("Title:", title);
      console.log("Description:", description);
      // Resetar o formulário ou navegar para outra tela
    }
  };

  return (
    <View style={{ flex: 1, alignItems: 'center', padding: window.height * 0.025 }}>
        <Text style={styles.title}>Você está registrando uma atividade para a sinergia entre:</Text>
        {cards.map((card, index) => (
          <CardSinergyProjects key={index} {...card} />
        ))}
      <Tabs step={step} setStep={setStep} page={'activity'}/>

      {step === 0 && (
        <View style={styles.dateContainer}>
          <Text style={{textAlign: 'left', fontWeight: 'bold', color: '#000', fontSize: 14}}>Selecione a data da atividade:</Text>
          <DateInput 
            placeholder="DD/MM/YYYY"
            value={date}
            onChangeText={setDate}
          />
        </View>
      )}

      {step === 1 && (
        <View style={styles.titleContainer}>
          <InputComponent
            type="input"
            label="Dê um título para o registro"
            placeholder={"Digite o título..."}
            maxLength={40}
          />
        </View>
      )}

      {step === 2 && (
        <View style={styles.descriptionContainer}>
          <InputComponent
            type="textarea"
            label="Adicione uma descrição"
            placeholder={"Digite a descrição..."}
            maxLength={2000}
          />
        </View>
      )}
      <View style={{ position: 'absolute', bottom: 0, left: 0, right: 0, marginBottom: 10 }}>
        <View style={{ flexDirection: 'row', alignSelf: 'center', width: window.width * 0.7 }}>
          <GradientButton onPress={handleNext} title='Continuar' />
        </View>
      </View>

    </View>
  );
}

const styles = StyleSheet.create({
  card: {
    padding: 10,
    margin: 5,
    backgroundColor: "#fff",
    borderRadius: 10,
    shadowColor: "#000",
    shadowOpacity: 0.1,
    shadowOffset: { width: 0, height: 2 },
    shadowRadius: 10,
    elevation: 5,
  },
  title: {
    fontSize: 14,
    fontWeight: 'bold',
    marginBottom: 5,
    textAlign: 'left',
  },
  tabsContainer: {
    flexDirection: 'row',
    justifyContent: 'space-around',
    marginVertical: 20,
    width: '100%',
    position: 'relative',
  },
  tab: {
    flex: 1,
    alignItems: 'center',
    marginBottom: window.height * 0.03,
  },
  tabText: {
    fontSize: 14,
  },
  activeTabText: {
    color: '#3A8A88',
    fontWeight: 'bold',
  },
  inactiveTabText: {
    color: '#999',
  },
  lineTab: {
    width: '90%',
    height: 6,
    borderRadius: 8,
    marginTop: 4,
  },
  activeLineTab: {
    backgroundColor: '#3A8A88',
  },
  inactiveLineTab: {
    backgroundColor: '#999',
  },
  dateContainer: {
    width: '100%',
    textAlign: 'left',
    fontWeight: 'bold',
  },
  titleContainer: {
    alignItems: 'center',
    width: '80%',
    fontWeight: 'bold', 
    color: '#999'
  },
  descriptionContainer: {
    alignItems: 'center',
    width: '80%',
  },
  input: {
    borderBottomWidth: 1,
    width: 200,
    marginBottom: 20,
  },
  continueButton: {
    backgroundColor: '#4CAF50',
    padding: 10,
    borderRadius: 5,
  },
  continueButtonText: {
    color: '#fff',
  },
});
