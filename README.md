# Preprocess 

I am starting this project because I'd like to have a useful tool to preprocess tabular datasets in a csv format (or any text format). 

I want a standalone tool that can do the following : 

### 🔹 1. **Data Cleaning**
- **Remove Duplicates**: Eliminate exact duplicate rows.
- **Handle Missing Values**:
  - Drop rows or columns with too many missing values.
  - Impute missing values using mean, median, mode, or predictive models.
- **Handle Outliers**:
  - Detect using IQR, Z-score, Isolation Forest, etc.
  - Remove or transform them.
- **Standardize Formats**: e.g., unify date formats, units, text casing.

---

### 🔹 2. **Encoding Categorical Variables**
- **Label Encoding**: Convert categories to integers.
- **One-Hot Encoding**: Create binary columns for each category.
- **Target Encoding / Mean Encoding**: Replace categories with the mean of the target variable.
- **Binary, Frequency Encoding**: Useful for high-cardinality categorical features.

---

### 🔹 3. **Feature Transformation**
- **Normalization (Min-Max Scaling)**: Scale features to a [0, 1] range.
- **Standardization (Z-score Scaling)**: Center to mean = 0, std = 1.
- **Log Transform, Box-Cox, Yeo-Johnson**: Reduce skewness in distributions.
- **Discretization (Binning)**: Convert continuous variables into categories.

---

### 🔹 4. **Dimensionality Reduction**
- **PCA (Principal Component Analysis)**.
- **t-SNE, UMAP**: Often for visualization.
- **Feature Selection**:
  - Based on feature importance.
  - Statistical tests (chi², ANOVA, correlation).
  - Regularization techniques (Lasso, Ridge).

---

### 🔹 5. **Feature Engineering**
- Create new features from existing ones (e.g., ratios, differences).
- Extract features from dates (day, month, weekday, etc.).
- Group-based aggregations (e.g., average purchase per user).
- Rolling statistics, lags (especially for time series).

---

### 🔹 6. **Handling Imbalanced Data**
- **Oversampling** (e.g., SMOTE, ADASYN).
- **Undersampling**.
- **Class Weighting**: Adjust loss function based on class frequency.

---

### 🔹 7. **Text Preprocessing (if applicable)**
- Clean text (remove punctuation, lowercase, remove stopwords).
- Tokenization.
- Lemmatization/Stemming.
- Vectorization (TF-IDF, CountVectorizer, embeddings like Word2Vec, BERT).

---

### 🔹 8. **Other Advanced Techniques**
- **Remove constant or quasi-constant columns**.
- **Handle time-series data**: create lag features, rolling windows, time-based splits.
- **Graph/sequential encodings** (for complex structured data).
