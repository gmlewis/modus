/*
 * This example is part of the Modus project, licensed under the Apache License 2.0.
 * You may modify and use this example in accordance with the license.
 * See the LICENSE file that accompanied this code for further details.
 */

import { models } from "@hypermode/modus-sdk-as";
import {
  ClassificationModel,
  ClassifierResult,
} from "@hypermode/modus-sdk-as/models/experimental/classification";

// This model name should match one defined in the modus.json manifest file.
const modelName: string = "my-classifier";

// This function takes input text and a probability threshold, and returns the
// classification label determined by the model, if the confidence is above the
// threshold. Otherwise, it returns an empty string.
export function classifyText(text: string, threshold: f32): string {
  const model = models.getModel<ClassificationModel>(modelName);
  const input = model.createInput([text]);
  const output = model.invoke(input);

  const prediction = output.predictions[0];
  if (prediction.confidence >= threshold) {
    return prediction.label;
  }

  return "";
}

// This function takes input text and returns the classification labels and their
// corresponding probabilities, as determined by the model.
export function getClassificationLabels(text: string): Map<string, f32> {
  const model = models.getModel<ClassificationModel>(modelName);
  const input = model.createInput([text]);
  const output = model.invoke(input);

  const prediction = output.predictions[0];
  const labels = getLabels(prediction);
  return labels;
}

function getLabels(prediction: ClassifierResult): Map<string, f32> {
  const labels = new Map<string, f32>();
  for (let i = 0; i < prediction.probabilities.length; i++) {
    const p = prediction.probabilities[i];
    labels.set(p.label, p.probability);
  }
  return labels;
}

// This function is similar to the previous, but allows multiple items to be classified at a time.
export function getMultipleClassificationLabels(
  ids: string[],
  texts: string[],
): Map<string, Map<string, f32>> {
  const model = models.getModel<ClassificationModel>(modelName);
  const input = model.createInput(texts);
  const output = model.invoke(input);

  const results = new Map<string, Map<string, f32>>();
  for (let i = 0; i < output.predictions.length; i++) {
    const labels = getLabels(output.predictions[i]);
    results.set(ids[i], labels);
  }
  return results;
}
